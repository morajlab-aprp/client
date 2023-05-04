from client_packages_logic import index


def test_index():
    assert index.hello() == "Hello client-packages-logic"
