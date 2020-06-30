echo "Creating channel..."
./scripts/createChannel.sh

echo "Joining Deliverer to channel..."
./scripts/join-peer.sh peer0 deliverer DelivererMSP 10051 1.0

echo "Joining Manufacturer to channel..."
./scripts/join-peer.sh peer0 manufacturer ManufacturerMSP 9051 1.0

echo "Joining Retailer to channel..."
./scripts/join-peer.sh peer0 retailer RetailerMSP 11051 1.0
        
echo "Installing chaincode for manufacturer..."
./scripts/install-chaincode.sh peer0 manufacturer ManufacturerMSP 9051 1.0

echo "Installing chaincode for deliverer..."
./scripts/install-chaincode.sh peer0 deliverer DelivererMSP 10051 1.0

echo "Installing chaincode for retailer..."
./scripts/install-chaincode.sh peer0 retailer RetailerMSP 11051 1.0

echo "Installing chaincode for producer..."
./scripts/install-chaincode.sh peer0 producer ProducerMSP 7051 1.0

echo "Instantiating chaincode on network..."
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
/usr/bin/time --format='la commande a duré %e secondes'  docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/users/Admin@manufacturer.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.example.com:9051" cli peer chaincode instantiate -n example -v 1.0 -C mychannel -c '{"Args":[]}' --tls --cafile $ORDERER_CA | ps aux | grep chaincode | grep -v grep | awk 'BEGIN { sum=0 } {sum=sum+$6; } END {printf("Taille RAM utilisée: %s Mo\n",sum / 1024)}' 
