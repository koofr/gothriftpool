typedef string UUID

struct MyRequest {
    1: required string req;
}

struct MyResult {
    1: required string res;
}

exception MyException {}

service MyService {

    void ping(),

    MyResult get_result(1:UUID id, 2:MyRequest req) throws (1:MyException me)

}
