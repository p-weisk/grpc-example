const grpc = require('grpc');
const PROTO_PATH = '../../api/api.proto';
const InvoiceService = grpc.load(PROTO_PATH).api.InvoiceService;
const client = new InvoiceService('localhost:7777',
    grpc.credentials.createInsecure());
module.exports = client;