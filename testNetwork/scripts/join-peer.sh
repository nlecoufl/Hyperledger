PEER=$1
ORG=$2
MSP=$3
PORT=$4
VERSION=$5

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_LOCALMSPID=$MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.example.com/peers/$PEER.$ORG.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.example.com/users/Admin@$ORG.example.com/msp
export CORE_PEER_ADDRESS=$PEER.$ORG.example.com:$PORT
export CHANNEL_NAME=mychannel
export CORE_PEER_TLS_ENABLED=true

docker exec -e "ORDERER_CA=$ORDERER_CA" -e "CORE_PEER_LOCALMSPID=$CORE_PEER_LOCALMSPID" -e "CORE_PEER_TLS_ROOTCERT_FILE=$CORE_PEER_TLS_ROOTCERT_FILE" -e "CORE_PEER_MSPCONFIGPATH=$CORE_PEER_MSPCONFIGPATH" -e "CORE_PEER_ADDRESS=$CORE_PEER_ADDRESS" -e "CHANNEL_NAME=$CHANNEL_NAME" -e "CORE_PEER_TLS_ENABLED=$CORE_PEER_TLS_ENABLED" cli peer channel join -b mychannel.block

