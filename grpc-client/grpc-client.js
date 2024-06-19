var PROTO_PATH = __dirname + '/../grpc-server/advice/advice.proto'

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var advice_proto = grpc.loadPackageDefinition(packageDefinition).advice;

function main() {
    var argv = parseArgs(process.argv.slice(2), {
        default: {
            target: 'localhost:50051'
        },
        string: 'target'
    });
    var target = argv.target;
    var client = new advice_proto.Advice(target,  grpc.credentials.createInsecure());
    client.GetAdvice({}, function(err, response) {
        if (!response) {
            console.error('Algo deu errado ao comunicar com o servidor. Verifique se ele está sendo executado.');
            return;
        }
        console.log('Advice:', response.message);
    });
}

main();