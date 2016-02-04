typedef string UUID

struct MyRequest {
    1: required string req;
}

struct MyResponse {
    1: required string res;
}

exception MyException {}

service MyService {

    void ping(),

    MyResponse get_response(1:UUID id, 2:MyRequest req) throws (1:MyException me)

}
