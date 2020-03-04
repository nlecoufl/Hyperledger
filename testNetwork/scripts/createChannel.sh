#!/bin/bash

echo "Creating channel..."
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_LOCALMSPID=ProducerMSP
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/users/Admin@producer.example.com/msp
export CORE_PEER_ADDRESS=peer0.producer.example.com:7051
export CHANNEL_NAME=mychannel
export CORE_PEER_TLS_ENABLED=true
export ORDERER_SYSCHAN_ID=syschain
docker exec -e "ORDERER_CA=$ORDERER_CA" -e "CORE_PEER_LOCALMSPID=$CORE_PEER_LOCALMSPID" -e "CORE_PEER_TLS_ROOTCERT_FILE=$CORE_PEER_TLS_ROOTCERT_FILE" -e "CORE_PEER_MSPCONFIGPATH=$CORE_PEER_MSPCONFIGPATH" -e "CORE_PEER_ADDRESS=$CORE_PEER_ADDRESS" -e "CHANNEL_NAME=$CHANNEL_NAME" -e "CORE_PEER_TLS_ENABLED=$CORE_PEER_TLS_ENABLED" -e "ORDERER_SYSCHAN_ID=$ORDERER_SYSCHAN_ID" cli peer channel create -o  orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
#peer channel create -o orderer.example.com:7050 /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -c mychannel -f ./channel-artifacts/channel.tx
echo 
echo "Channel created, joining Producer..."
docker exec -e "ORDERER_CA=$ORDERER_CA" -e "CORE_PEER_LOCALMSPID=$CORE_PEER_LOCALMSPID" -e "CORE_PEER_TLS_ROOTCERT_FILE=$CORE_PEER_TLS_ROOTCERT_FILE" -e "CORE_PEER_MSPCONFIGPATH=$CORE_PEER_MSPCONFIGPATH" -e "CORE_PEER_ADDRESS=$CORE_PEER_ADDRESS" -e "CHANNEL_NAME=$CHANNEL_NAME" -e "CORE_PEER_TLS_ENABLED=$CORE_PEER_TLS_ENABLED" -e "ORDERER_SYSCHAN_ID=$ORDERER_SYSCHAN_ID" cli peer channel join -b mychannel.block
