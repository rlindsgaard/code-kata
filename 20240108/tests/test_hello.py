from hello import Hello
from pytest import fixture

@fixture
def hello():
    return Hello()

def test_hello(hello):
    assert hello.message == "Hello from python3 class!"