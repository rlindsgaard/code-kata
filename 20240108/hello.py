#!/usr/bin/env python3

class Hello:
    def __init__(self):
        self.message = "Hello from python3 class!"

    def say_hello(self):
        print(self.message)

if __name__ == "__main__":
    hello = Hello()
    hello.say_hello()   