import re
filepath = "/Users/0x1/Documents/htdocs/gopurs/gopurs-node-buffer/src/Node/Buffer/Immutable.go"
with open(filepath, 'r') as f:
    content = f.read()

content = content.replace("getBytes", "nodeBufferImmutable_getBytes")
content = content.replace("boxBytes", "nodeBufferImmutable_boxBytes")

with open(filepath, 'w') as f:
    f.write(content)
print("Replaced getBytes and boxBytes in Immutable.go")
