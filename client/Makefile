PROTO_DIR=app
PROTO_FILE=user.proto
PYTHON_DIR=app
GRPC_TOOLS=python -m grpc_tools.protoc

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make proto        - Generate gRPC code from proto files"
	@echo "  make run          - Run the FastAPI application"
	@echo "  make clean        - Clean up generated files"
	@echo "  make install      - Install dependencies using Poetry"
	@echo "  make test         - Run tests"

.PHONY: proto
proto:
	$(GRPC_TOOLS) -I$(PROTO_DIR) --python_out=$(PYTHON_DIR) --grpc_python_out=$(PYTHON_DIR) $(PROTO_DIR)/$(PROTO_FILE)
	sed -i 's/^import user_pb2 as/from . import user_pb2 as/' $(PYTHON_DIR)/user_pb2_grpc.py

.PHONY: run
run:
	python -m venv .env
	source .env/bin/activate
	make install
	poetry run uvicorn app.main:app --host 0.0.0.0 --port 8000

.PHONY: install
install:
	poetry install

.PHONY: clean
clean:
	rm -f $(PYTHON_DIR)/$(PROTO_FILE).py
	rm -f $(PYTHON_DIR)/$(PROTO_FILE)_pb2.py
	rm -f $(PYTHON_DIR)/$(PROTO_FILE)_pb2_grpc.py
	find . -type d -name '__pycache__' -exec rm -r {} +

.PHONY: test
test:
	poetry run pytest
