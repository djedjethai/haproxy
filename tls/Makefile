CONFIG_PATH=$(shell pwd)/certifs
CONFIG_GPT=$(shell pwd)/certifgpt


.PHONY: gencert
gencert:
	cfssl gencert \
		-initca test/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		test/client-csr.json | cfssljson -bare client	
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=server \
		test/server-csr.json | cfssljson -bare server
	mv *.pem *.csr ${CONFIG_PATH}

# does not work...
.PHONY: gengpt
gengpt:
	cfssl gencert \
		-initca testgpt/ca-csr.json | cfssljson -bare ca
	cfssl genkey \
		-config testgpt/server-config.json \
		-profile server \
		testgpt/server-csr.json | cfssljson -bare server
	cfssl sign -ca ca.pem -ca-key ca-key.pem -
	mv *.pem *.csr ${CONFIG_GPT}
	


	

