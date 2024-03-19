import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';


const client = new Client();
client.load(['../../api/demo'], 'demo.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';

export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });
    const data = {
        uid: 'test',
        message: 'hello',
    };

    let response = client.invoke('demo.pb.DemoService/Hi', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    client.close();
    sleep(1);
}


