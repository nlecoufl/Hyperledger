var express = require('express');
var bodyParser = require('body-parser');

var app = express();
app.use(bodyParser.json());

// Setting for Hyperledger Fabric
const { FileSystemWallet, Gateway } = require('fabric-network');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..','..', 'connections', 'connection-manufacturer.json');


app.get('/api/queryalldatas', async function (req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), '../wallet/wallet-manufacturer/');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('example');

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction('queryAllDatas');
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        res.status(200).json({response: result.toString()});

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({error: error});
        process.exit(1);
    }
});


app.get('/api/query/:data_index', async function (req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), '../wallet/wallet-manufacturer/');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('example');

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction('queryData', req.params.data_index);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        res.status(200).json({response: result.toString()});

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({error: error});
        process.exit(1);
    }
});

app.post('/api/addbulkdatas/', async function (req, res) {
    try {
  
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), '../wallet/wallet-manufacturer/');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
  
        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });
        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');
  
        // Get the contract from the network.
        const contract = network.getContract('example');
  
        // Submit the specified transaction.
        // createCar transaction - requires 5 argument, ex: ('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom')
        // changeCarOwner transaction - requires 2 args , ex: ('changeCarOwner', 'CAR10', 'Dave')
        const json = '{"lastUpdatedOther":1579850797,"ttl":3600,"data":{"stations":[{"stationCode":"16107","station_id":213688169,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":33,"numDocksAvailable":33,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850732},{"stationCode":"6015","station_id":99950133,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":6},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850255},{"stationCode":"9020","station_id":36255,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850484},{"stationCode":"12109","station_id":37815204,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":26,"numDocksAvailable":26,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850753},{"stationCode":"5001","station_id":100769544,"num_bikes_available":8,"numBikesAvailable":8,"num_bikes_available_types":[{"mechanical":4},{"ebike":4}],"num_docks_available":30,"numDocksAvailable":30,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850786},{"stationCode":"14014","station_id":85002689,"num_bikes_available":11,"numBikesAvailable":11,"num_bikes_available_types":[{"mechanical":9},{"ebike":2}],"num_docks_available":47,"numDocksAvailable":47,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850734},{"stationCode":"17026","station_id":54000559,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":36,"numDocksAvailable":36,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850765},{"stationCode":"17041","station_id":85043758,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":35,"numDocksAvailable":35,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850729},{"stationCode":"10013","station_id":123095125,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":56,"numDocksAvailable":56,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850678},{"stationCode":"9104","station_id":129026597,"num_bikes_available":7,"numBikesAvailable":7,"num_bikes_available_types":[{"mechanical":4},{"ebike":3}],"num_docks_available":15,"numDocksAvailable":15,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850421},{"stationCode":"14111","station_id":251039991,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":2},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850746},{"stationCode":"6003","station_id":37874517,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850219},{"stationCode":"13007","station_id":66491398,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":45,"numDocksAvailable":45,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850720},{"stationCode":"5110","station_id":210403080,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850689},{"stationCode":"6108","station_id":210561800,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":4},{"ebike":0}],"num_docks_available":13,"numDocksAvailable":13,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850370},{"stationCode":"33006","station_id":209063434,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":2},{"ebike":7}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850316},{"stationCode":"42016","station_id":94555589,"num_bikes_available":3,"numBikesAvailable":3,"num_bikes_available_types":[{"mechanical":1},{"ebike":2}],"num_docks_available":24,"numDocksAvailable":24,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850779},{"stationCode":"41301","station_id":244498842,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850627},{"stationCode":"10105","station_id":210738704,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":25,"numDocksAvailable":25,"is_installed":1,"is_returning":0,"is_renting":0,"last_reported":1579850576},{"stationCode":"21010","station_id":43195240,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850515}]}}'
        const obj = JSON.parse(json);
        objectLength = Object.keys(obj.data.stations).length
        for (let i = 0; i < objectLength; i++) {
          var pas = i+21;
          str = "DATA" + pas;  
          await contract.submitTransaction('createDatas', str, obj.data.stations[i].stationCode, obj.data.stations[i].station_id.toString(), obj.data.stations[i].num_bikes_available.toString(), obj.data.stations[i].numBikesAvailable.toString(), obj.data.stations[i].num_bikes_available_types[0].mechanical.toString(), obj.data.stations[i].num_bikes_available_types[1].ebike.toString(), obj.data.stations[i].num_docks_available.toString(), obj.data.stations[i].numDocksAvailable.toString(), obj.data.stations[i].is_installed.toString(), obj.data.stations[i].is_returning.toString(), obj.data.stations[i].is_renting.toString(), obj.data.stations[i].last_reported.toString());
          console.log('Transaction has been submitted');
        }
        res.send('Transaction has been submitted');
        // Disconnect from the gateway.
        await gateway.disconnect();
  
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
});


app.listen(8080);

