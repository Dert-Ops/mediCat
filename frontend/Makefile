all:
	@echo "If your first run, please run 'make init' first."
	@echo "Run 'make run' to start the server."
	@echo "Run 'make dev' to start the server in developer mode."

init:
	@echo "Install dependencies..."
	@npm install
	@echo "Build..."
	@npm run build

dev:
	@echo "Start server developer mode..."
	@npm run dev

run:
	@echo "Start server..."
	@npm run preview

clean:
	@echo "Clean..."
	@rm -rf node_modules
	@rm -rf dist

.PHONY: all init run clean