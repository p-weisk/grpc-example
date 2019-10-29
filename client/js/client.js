const grpc = require('grpc');
const PROTO_PATH = '../../api/api.proto';
const InvoiceService = grpc.load(PROTO_PATH).api.InvoiceService;
const ProductService = grpc.load(PROTO_PATH).api.ProductService;
const invoiceClient = new InvoiceService('localhost:7777',
    grpc.credentials.createInsecure());
const productClient = new ProductService('localhost:7777',
    grpc.credentials.createInsecure())
module.exports = { invoiceClient: invoiceClient, productClient: productClient };