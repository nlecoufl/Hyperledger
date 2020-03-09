# **Hyperledger**
Hyperledger testing

## **Windows**
### **Prerequis**
* Docker  https://store.docker.com/editions/community/docker-ce-desktop-windows
* Go      https://golang.org/doc/install?download=go1.11.2.windows-amd64.msi
* Node.js https://nodejs.org/en/download/
* Git     https://git-scm.com/download/win

Suivre : https://www.linkedin.com/pulse/step-hyperledger-fabric-installation-windows-haris-mumtaz/

Activer Hyper-V : https://docs.microsoft.com/fr-fr/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v?redirectedfrom=MSDN

### **Set up**
#### Etape 1 : Obtenir l'application et les outils Hyperledger Fabric
Dans cette démo, on utilise Hyperledger Fabric v1.4.4.
Vérifier que l'application et les outils ont été installés en regardant :
* le dossier fabric-samples, dans lequel les outils sont dans le dossier /bin
* les images docker ($ docker images hyperledger/fabric*)

Sinon : 
        mkdir hyperledger
        cd hyperledger
        curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.4

#### Etape 2 : Design notre network
Architecture de notre réseau Hyperledger :
* une organisation Orderer (OrdererOrg) avec un noeud Orderer
* quatre organisations (Producer, Manufacturer, Deliverer, Retailer) avec un peer pour chaque organisation
* un channel pour toutes les organisations

On créé notre dossier :
    
    cd hyperledger
    mkdir Hyperledger_v0
    ls
    > fabric-samples Hyperledger_v0
    cd Hyperledger_v0 
    mkdir testNetwork && cd testNetwork


#### Etape 2.1 : Fichier de configuration crypto-config.yaml
Il permet de générer les certificats pour les différentes parties, dont les noeuds et les administrateurs.

    cd Hyperledger_v0
    touch crypto-config.yaml

Crypto-config.yaml :

    OrdererOrgs:
    - Name: Orderer
        Domain: example.com
        EnableNodeOUs: true
        Specs:
        - Hostname: orderer

    PeerOrgs:
    - Name: Producer
        Domain: producer.example.com
        EnableNodeOUs: true
        Template:
        Count: 1
        Users:
        Count: 1

    - Name: Manufacturer
        Domain: manufacturer.example.com
        EnableNodeOUs: true
        Template:
        Count: 1
        Users:
        Count: 1

    - Name: Deliverer
        Domain: deliverer.example.com
        EnableNodeOUs: true
        Template:
        Count: 1
        Users:
        Count: 1

    - Name: Retailer
        Domain: retailer.example.com
        EnableNodeOUs: true
        Template:
        Count: 1
        Users:
        Count: 1

#### Etape 2.2 : Fichier configtx.yaml
Ce fichier conserve tout le réseau d'entreprises que nous construisons. C'est ainsi que le réseau est réellement construit. Parmi les nombreuses configurations par défaut, nous nous concentrons sur les Profils (la partie inférieure de ce fichier). Les profils sont utilisés pour générer le bloc de genèse et les transactions nécessaires à la mise en place du réseau

Le profil du bloc genesis est OrdererGenesis, et trois organisations y sont incluses.
Le profil pour le channel est ChannelAll.

    cd Hyperledger_v0
    touch configtx.yaml

Fichier configtx.yaml
    
    ---
    Organizations:
        - &OrdererOrg
            Name: OrdererOrg
            ID: OrdererMSP
            MSPDir: crypto-config/ordererOrganizations/example.com/msp
            Policies:
                Readers:
                    Type: Signature
                    Rule: "OR('OrdererMSP.member')"
                Writers:
                    Type: Signature
                    Rule: "OR('OrdererMSP.member')"
                Admins:
                    Type: Signature
                    Rule: "OR('OrdererMSP.admin')"

        - &Producer
            Name: ProducerMSP
            ID: ProducerMSP
            MSPDir: crypto-config/peerOrganizations/producer.example.com/msp
            Policies:
                Readers:
                    Type: Signature
                    Rule: "OR('ProducerMSP.admin', 'ProducerMSP.peer', 'ProducerMSP.client')"
                Writers:
                    Type: Signature
                    Rule: "OR('ProducerMSP.admin', 'ProducerMSP.client')"
                Admins:
                    Type: Signature
                    Rule: "OR('ProducerMSP.admin')"

        - &Manufacturer
            Name: ManufacturerMSP
            ID: ManufacturerMSP
            MSPDir: crypto-config/peerOrganizations/manufacturer.example.com/msp
            Policies:
                Readers:
                    Type: Signature
                    Rule: "OR('ManufacturerMSP.admin', 'ManufacturerMSP.peer', 'ManufacturerMSP.client')"
                Writers:
                    Type: Signature
                    Rule: "OR('ManufacturerMSP.admin', 'ManufacturerMSP.client')"
                Admins:
                    Type: Signature
                    Rule: "OR('ManufacturerMSP.admin')"
                    
        - &Deliverer
            Name: DelivererMSP
            ID: DelivererMSP
            MSPDir: crypto-config/peerOrganizations/deliverer.example.com/msp
            Policies:
                Readers:
                    Type: Signature
                    Rule: "OR('DelivererMSP.admin', 'DelivererMSP.peer', 'DelivererMSP.client')"
                Writers:
                    Type: Signature
                    Rule: "OR('DelivererMSP.admin', 'DelivererMSP.client')"
                Admins:
                    Type: Signature
                    Rule: "OR('DelivererMSP.admin')"
                    
        - &Retailer
            Name: RetailerMSP
            ID: RetailerMSP
            MSPDir: crypto-config/peerOrganizations/retailer.example.com/msp
            Policies:
                Readers:
                    Type: Signature
                    Rule: "OR('RetailerMSP.admin', 'RetailerMSP.peer', 'RetailerMSP.client')"
                Writers:
                    Type: Signature
                    Rule: "OR('RetailerMSP.admin', 'RetailerMSP.client')"
                Admins:
                    Type: Signature
                    Rule: "OR('RetailerMSP.admin')"

    Capabilities:
        Channel: &ChannelCapabilities
            V1_4_3: true
            V1_3: false
            V1_1: false
        Orderer: &OrdererCapabilities
            V1_4_2: true
            V1_1: false
        Application: &ApplicationCapabilities
            V1_4_2: true
            V1_3: false
            V1_2: false
            V1_1: false

    Application: &ApplicationDefaults
        Organizations:
        Policies:
            Readers:
                Type: ImplicitMeta
                Rule: "ANY Readers"
            Writers:
                Type: ImplicitMeta
                Rule: "ANY Writers"
            Admins:
                Type: ImplicitMeta
                Rule: "MAJORITY Admins"

        Capabilities:
            <<: *ApplicationCapabilities

    Orderer: &OrdererDefaults
        OrdererType: solo
        Addresses:
            - orderer.example.com:7050
        BatchTimeout: 2s
        BatchSize:
            MaxMessageCount: 10
            AbsoluteMaxBytes: 99 MB
            PreferredMaxBytes: 512 KB

        Kafka:
            Brokers:
                - 127.0.0.1:9092

    Organizations:
        Policies:
            Readers:
                Type: ImplicitMeta
                Rule: "ANY Readers"
            Writers:
                Type: ImplicitMeta
                Rule: "ANY Writers"
            Admins:
                Type: ImplicitMeta
                Rule: "MAJORITY Admins"
            BlockValidation:
                Type: ImplicitMeta
                Rule: "ANY Writers"

    Channel: &ChannelDefaults
        Policies:
            Readers:
                Type: ImplicitMeta
                Rule: "ANY Readers"
            Writers:
                Type: ImplicitMeta
                Rule: "ANY Writers"
            Admins:
                Type: ImplicitMeta
                Rule: "MAJORITY Admins"
        Capabilities:
            <<: *ChannelCapabilities

    Profiles:
        OrdererGenesis:
            <<: *ChannelDefaults
            Orderer:
                <<: *OrdererDefaults
                Organizations:
                    - *OrdererOrg
                Capabilities:
                    <<: *OrdererCapabilities
            Consortiums:
                SampleConsortium:
                    Organizations:
                        - *Producer
                        - *Manufacturer
                        - *Deliverer
                        - *Retailer
        ChannelAll:
            Consortium: SampleConsortium
            <<: *ChannelDefaults
            Application:
                <<: *ApplicationDefaults
                Organizations:
                    - *Producer
                    - *Manufacturer
                    - *Deliverer
                    - *Retailer
                Capabilities:
                    <<: *ApplicationCapabilities

#### Etape 2.3 : Fichiers docker-compose
Fichier docker-compose.yaml :

    version: '2'

    volumes:
    orderer.example.com:
    peer0.producer.example.com:
    peer0.manufacturer.example.com:
    peer0.deliverer.example.com:
    peer0.retailer.example.com:
    caProducer:
    caManufacturer:
    caDeliverer:
    caRetailer:

    networks:
    byfn:

    services:

    orderer.example.com:
        extends:
        file:   base/docker-compose-base.yaml
        service: orderer.example.com
        container_name: orderer.example.com
        networks:
        - byfn

    peer0.producer.example.com:
        container_name: peer0.producer.example.com
        extends:
        file:  base/docker-compose-base.yaml
        service: peer0.producer.example.com
        networks:
        - byfn

    peer0.manufacturer.example.com:
        container_name: peer0.manufacturer.example.com
        extends:
        file:  base/docker-compose-base.yaml
        service: peer0.manufacturer.example.com
        networks:
        - byfn
        
    peer0.deliverer.example.com:
        container_name: peer0.deliverer.example.com
        extends:
        file:  base/docker-compose-base.yaml
        service: peer0.deliverer.example.com
        networks:
        - byfn
        
    peer0.retailer.example.com:
        container_name: peer0.retailer.example.com
        extends:
        file:  base/docker-compose-base.yaml
        service: peer0.retailer.example.com
        networks:
        - byfn
        
    caProducer:
        container_name: caProducer
        extends:
        file:  base/docker-compose-base.yaml
        service: caProducer
        networks:
        - byfn
        
    caManufacturer:
        container_name: caManufacturer
        extends:
        file:  base/docker-compose-base.yaml
        service: caManufacturer
        networks:
        - byfn
        
    caDeliverer:
        container_name: caDeliverer
        extends:
        file:  base/docker-compose-base.yaml
        service: caDeliverer
        networks:
        - byfn
        
    caRetailer:
        container_name: caRetailer
        extends:
        file:  base/docker-compose-base.yaml
        service: caRetailer
        networks:
        - byfn
        
        
    cli:
        container_name: cli
        image: hyperledger/fabric-tools:$IMAGE_TAG
        tty: true
        stdin_open: true
        environment:
        - GOPATH=/opt/gopath
        - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
        #- FABRIC_LOGGING_SPEC=DEBUG
        - FABRIC_LOGGING_SPEC=INFO
        - CORE_PEER_ID=cli
        - CORE_PEER_ADDRESS=peer0.producer.example.com:7051
        - CORE_PEER_LOCALMSPID=ProducerMSP
        - CORE_PEER_TLS_ENABLED=true
        - CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/server.crt
        - CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/server.key
        - CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls/ca.crt
        - CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/producer.example.com/users/Admin@producer.example.com/msp
        working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
        command: /bin/bash
        volumes:
            - /var/run/:/host/var/run/
            - ./../chaincode/:/opt/gopath/src/github.com/chaincode
            - ./crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/
            - ./scripts:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/
            - ./channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts
        depends_on:
        - orderer.example.com
        - peer0.producer.example.com
        - peer0.manufacturer.example.com
        - peer0.deliverer.example.com
        - peer0.retailer.example.com
        - caProducer
        - caManufacturer
        - caDeliverer
        - caRetailer
        networks:
        - byfn

Fichier base/docker-compose-base.yaml :

    version: '2'

    services:

    orderer.example.com:
        container_name: orderer.example.com
        extends:
        file: peer-base.yaml
        service: orderer-base
        volumes:
            - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
            - ../crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp:/var/hyperledger/orderer/msp
            - ../crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/:/var/hyperledger/orderer/tls
            - orderer.example.com:/var/hyperledger/production/orderer
        ports:
        - 7050:7050

    peer0.producer.example.com:
        container_name: peer0.producer.example.com
        extends:
        file: peer-base.yaml
        service: peer-base
        environment:
        - CORE_PEER_ID=peer0.producer.example.com
        - CORE_PEER_ADDRESS=peer0.producer.example.com:7051
        - CORE_PEER_LISTENADDRESS=0.0.0.0:7051
        - CORE_PEER_CHAINCODEADDRESS=peer0.producer.example.com:7052
        - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052
        - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.manufacturer.example.com:7051
        - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.producer.example.com:7051
        - CORE_PEER_LOCALMSPID=ProducerMSP
        volumes:
            - /var/run/:/host/var/run/
            - ../crypto-config/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/msp:/etc/hyperledger/fabric/msp
            - ../crypto-config/peerOrganizations/producer.example.com/peers/peer0.producer.example.com/tls:/etc/hyperledger/fabric/tls
            - peer0.producer.example.com:/var/hyperledger/production
        ports:
        - 7051:7051

    peer0.manufacturer.example.com:
        container_name: peer0.manufacturer.example.com
        extends:
        file: peer-base.yaml
        service: peer-base
        environment:
        - CORE_PEER_ID=peer0.manufacturer.example.com
        - CORE_PEER_ADDRESS=peer0.manufacturer.example.com:9051
        - CORE_PEER_LISTENADDRESS=0.0.0.0:9051
        - CORE_PEER_CHAINCODEADDRESS=peer0.manufacturer.example.com:9052
        - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:9052
        - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.manufacturer.example.com:9051
        - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.producer.example.com:9051
        - CORE_PEER_LOCALMSPID=ManufacturerMSP
        volumes:
            - /var/run/:/host/var/run/
            - ../crypto-config/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/msp:/etc/hyperledger/fabric/msp
            - ../crypto-config/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/tls:/etc/hyperledger/fabric/tls
            - peer0.manufacturer.example.com:/var/hyperledger/production
        ports:
        - 9051:9051
        
    peer0.deliverer.example.com:
        container_name: peer0.deliverer.example.com
        extends:
        file: peer-base.yaml
        service: peer-base
        environment:
        - CORE_PEER_ID=peer0.deliverer.example.com
        - CORE_PEER_ADDRESS=peer0.deliverer.example.com:10051
        - CORE_PEER_LISTENADDRESS=0.0.0.0:10051
        - CORE_PEER_CHAINCODEADDRESS=peer0.deliverer.example.com:10052
        - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:10052
        - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.deliverer.example.com:10051
        - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.producer.example.com:10051
        - CORE_PEER_LOCALMSPID=DelivererMSP
        volumes:
            - /var/run/:/host/var/run/
            - ../crypto-config/peerOrganizations/deliverer.example.com/peers/peer0.deliverer.example.com/msp:/etc/hyperledger/fabric/msp
            - ../crypto-config/peerOrganizations/deliverer.example.com/peers/peer0.deliverer.example.com/tls:/etc/hyperledger/fabric/tls
            - peer0.deliverer.example.com:/var/hyperledger/production
        ports:
        - 10051:10051
        
        
    peer0.retailer.example.com:
        container_name: peer0.retailer.example.com
        extends:
        file: peer-base.yaml
        service: peer-base
        environment:
        - CORE_PEER_ID=peer0.retailer.example.com
        - CORE_PEER_ADDRESS=peer0.retailer.example.com:11051
        - CORE_PEER_LISTENADDRESS=0.0.0.0:11051
        - CORE_PEER_CHAINCODEADDRESS=peer0.retailer.example.com:11052
        - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:11052
        - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.retailer.example.com:11051
        - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.producer.example.com:11051
        - CORE_PEER_LOCALMSPID=RetailerMSP
        volumes:
            - /var/run/:/host/var/run/
            - ../crypto-config/peerOrganizations/retailer.example.com/peers/peer0.retailer.example.com/msp:/etc/hyperledger/fabric/msp
            - ../crypto-config/peerOrganizations/retailer.example.com/peers/peer0.retailer.example.com/tls:/etc/hyperledger/fabric/tls
            - peer0.retailer.example.com:/var/hyperledger/production
        ports:
        - 11051:11051
        
    caProducer:
        image: hyperledger/fabric-ca:$IMAGE_TAG
        environment:
        - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
        - FABRIC_CA_SERVER_CA_NAME=ca-producer
        - FABRIC_CA_SERVER_TLS_ENABLED=true
        - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.producer.example.com-cert.pem
        - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/8cabfa8f02441dbea56a332b1331c77842bf7ea0120bfe5a611020811f906dc1_sk
        - FABRIC_CA_SERVER_PORT=7054
        ports:
        - "7054:7054"
        command: sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.producer.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/8cabfa8f02441dbea56a332b1331c77842bf7ea0120bfe5a611020811f906dc1_sk -b admin:adminpw -d'
        volumes:
        - ../crypto-config/peerOrganizations/producer.example.com/ca/:/etc/hyperledger/fabric-ca-server-config
        
    caManufacturer:
        image: hyperledger/fabric-ca:$IMAGE_TAG
        environment:
        - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
        - FABRIC_CA_SERVER_CA_NAME=ca-manufacturer
        - FABRIC_CA_SERVER_TLS_ENABLED=true
        - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.manufacturer.example.com-cert.pem
        - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/8dcafcec62719d250863d21412c81321a8be611bee84600b7c71010e2f8e7620_sk
        - FABRIC_CA_SERVER_PORT=8054
        ports:
        - "8054:8054"
        command: sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.manufacturer.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/8dcafcec62719d250863d21412c81321a8be611bee84600b7c71010e2f8e7620_sk -b admin:adminpw -d'
        volumes:
        - ../crypto-config/peerOrganizations/manufacturer.example.com/ca/:/etc/hyperledger/fabric-ca-server-config
        
    caDeliverer:
        image: hyperledger/fabric-ca:$IMAGE_TAG
        environment:
        - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
        - FABRIC_CA_SERVER_CA_NAME=ca-deliverer
        - FABRIC_CA_SERVER_TLS_ENABLED=true
        - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.deliverer.example.com-cert.pem
        - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/e2fc8bdf0f57f98ef83505378b1e270ef465d08e1a6bed0adae1aa61a665b19c_sk
        - FABRIC_CA_SERVER_PORT=9054
        ports:
        - "9054:9054"
        command: sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.deliverer.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/e2fc8bdf0f57f98ef83505378b1e270ef465d08e1a6bed0adae1aa61a665b19c_sk -b admin:adminpw -d'
        volumes:
        - ../crypto-config/peerOrganizations/deliverer.example.com/ca/:/etc/hyperledger/fabric-ca-server-config
        
        
    caRetailer:
        image: hyperledger/fabric-ca:$IMAGE_TAG
        environment:
        - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
        - FABRIC_CA_SERVER_CA_NAME=ca-retailer
        - FABRIC_CA_SERVER_TLS_ENABLED=true
        - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.retailer.example.com-cert.pem
        - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/5edb1fb62ee4a7b72ba2442bebc5dca70420cdbfc13cfa3c9a964125defc66cf_sk
        - FABRIC_CA_SERVER_PORT=10054
        ports:
        - "10054:10054"
        command: sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.retailer.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/5edb1fb62ee4a7b72ba2442bebc5dca70420cdbfc13cfa3c9a964125defc66cf_sk -b admin:adminpw -d'
        volumes:
        - ../crypto-config/peerOrganizations/retailer.example.com/ca/:/etc/hyperledger/fabric-ca-server-config

Fichier peer-base.yaml :

    version: '2'

    services:
    peer-base:
        image: hyperledger/fabric-peer:$IMAGE_TAG
        environment:
        - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
        # the following setting starts chaincode containers on the same
        # bridge network as the peers
        # https://docs.docker.com/compose/networking/
        - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=net_byfn
        - FABRIC_LOGGING_SPEC=INFO
        #- FABRIC_LOGGING_SPEC=DEBUG
        - CORE_PEER_TLS_ENABLED=true
        - CORE_PEER_GOSSIP_USELEADERELECTION=true
        - CORE_PEER_GOSSIP_ORGLEADER=false
        - CORE_PEER_PROFILE_ENABLED=true
        - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
        - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
        - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
        working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
        command: peer node start

    orderer-base:
        image: hyperledger/fabric-orderer:$IMAGE_TAG
        environment:
        - FABRIC_LOGGING_SPEC=INFO
        - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
        - ORDERER_GENERAL_GENESISMETHOD=file
        - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
        - ORDERER_GENERAL_LOCALMSPID=OrdererMSP
        - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
        # enabled TLS
        - ORDERER_GENERAL_TLS_ENABLED=true
        - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
        - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
        - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
        - ORDERER_KAFKA_TOPIC_REPLICATIONFACTOR=1
        - ORDERER_KAFKA_VERBOSE=true
        - ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt
        - ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key
        - ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
        working_dir: /opt/gopath/src/github.com/hyperledger/fabric
        command: orderer

#### Etape 2.4 : Préparer fichier .env
Préparer le .env contenant les variables utilisées dans docker-compose.yaml :

    echo COMPOSE_PROJECT_NAME=net >> .env
    echo IMAGE_TAG=latest >> .env

#### Etape 3 : Génération Certificats et Configuration des transactions
À l'étape 2, nous avons préparé les fichiers de configuration. Les fichiers crypto-config.yaml et configtx.yaml sont utilisés pour générer les fichiers de configuration qui seront utilisés lors de la mise en place de notre réseau.
Les deux outils sont stockés dans fabric-samples/bin. Ils sont respectivement cryptogen et configtxgen.

#### Etape 3.1 : Générer certificats pour les organisations, administrateurs et utilisateurs
Les certificats sont générés en utilisant l'outil cryptogen avec le fichier de configuration crypto-config.yaml. Cela créé un dossier crypto-config, contenant tous les certificats.
    ../bin/cryptogen generate --config=./crypto-config.yaml

#### Etape 3.2 : Générer les configurations des transactions pour les channels
L'outil que nous utilisons est configtxgen et le fichier de configuration est configtx.yaml. Notez que nous devons créer un répertoire channel-artifacts car les résultats sont stockés dans ce répertoire. Vous verrez une erreur si vous n'avez pas ce répertoire.

D'abord, nous générons le bloc de genèse en utilisant le profil OrdererGenesis. Le résultat est un fichier appelé genesis.block.

    mkdir channel-artifacts
    ../bin/configtxgen -profile OrdererGenesis -outputBlock ./channel-artifacts/genesis.block

Ensuite, nous générons la channel transaction.

    ../bin/configtxgen -profile ChannelAll -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel

#### Etape 4 : Lancer l'infrastructure
    docker-compose up -d
Les étapes précédentes permettent de comprendre le fonctionnement du réseau, afin de simplifier la suite, des scripts ont été créés.
    ./scripts/main.sh

Il permet de créer un channel entre toutes les organisations, d'installer un chaincode puis de l'instantier.
Fichier /scripts/main.sh :

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

    echo "Instantiating chaincode on network..."
    export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/users/Admin@manufacturer.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.example.com:9051" cli peer chaincode instantiate -n example -v 1.0 -C mychannel -c '{"Args":[]}' --tls --cafile $ORDERER_CA
    echo "Invoke initLedger on network..."
    sleep 10s # Waits 5 seconds.
    docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/tls/ca.crt" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/users/Admin@manufacturer.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.example.com:9051" cli peer chaincode invoke -n example  -C mychannel -c '{"Args":["initLedger"]}' --tls --cafile $ORDERER_CA

#### Etape 5 : Api
Le dossier src/ contient le server permettant de requêter sur notre chaincode.

        cd /src/server/
        node serverbis.js

Dans un autre terminal, on peut essayer :

        curl http://localhost:8080/api/queryalldatas
        curl http://localhost:8080/api/query/DATA4
        curl -H "Content-Type: application/json" -X POST http://localhost:8080/api/initbulk
        curl http://localhost:8080/api/query/DATA1300
