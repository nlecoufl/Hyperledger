echo "CA1_PRIVATE_KEY=$(cd crypto-config/peerOrganizations/producer.example.com/ca && ls *_sk)" >> .env
echo "CA2_PRIVATE_KEY=$(cd crypto-config/peerOrganizations/manufacturer.example.com/ca && ls *_sk)" >> .env
echo "CA3_PRIVATE_KEY=$(cd crypto-config/peerOrganizations/deliverer.example.com/ca && ls *_sk)" >> .env
echo "CA4_PRIVATE_KEY=$(cd crypto-config/peerOrganizations/retailer.example.com/ca && ls *_sk)" >> .env

