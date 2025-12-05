.PHONY: go-d1p1 go-d1p2 python-d1p1 python-d1p2 ts-d1p1 ts-d1p2

# Color codes
BOLD := \033[1m
RESET := \033[0m
CYAN := \033[36m
BLUE := \033[34m
YELLOW := \033[33m

define print_py_title
	@printf "\n$(YELLOW)$(BOLD)────────────────────────────────────────────$(RESET)\n"
	@printf "$(YELLOW)$(BOLD)🎄 Advent of Code 2025 | Python | $(1)$(RESET)\n"
	@printf "$(YELLOW)$(BOLD)────────────────────────────────────────────$(RESET)\n"
endef

define print_go_title
	@printf "\n$(CYAN)$(BOLD)────────────────────────────────────────────$(RESET)\n"
	@printf "$(CYAN)$(BOLD)🎄 Advent of Code 2025 | Go | $(1)$(RESET)\n"
	@printf "$(CYAN)$(BOLD)────────────────────────────────────────────$(RESET)\n"
endef

define print_ts_title
	@printf "\n$(BLUE)$(BOLD)────────────────────────────────────────────$(RESET)\n"
	@printf "$(BLUE)$(BOLD)🎄 Advent of Code 2025 | TypeScript | $(1)$(RESET)\n"
	@printf "$(BLUE)$(BOLD)────────────────────────────────────────────$(RESET)\n"
endef

# ------------- PYTHON -------------
# Day 1
python-d1p1:
	$(call print_py_title,Day 1 Part 1)
	@cd 2025/python && python main.py d1p1

python-d1p2:
	$(call print_py_title,Day 1 Part 2)
	@cd 2025/python && python main.py d1p2

# Day 2
python-d2p1:
	$(call print_py_title,Day 2 Part 1)
	@cd 2025/python && python main.py d2p1

python-d2p2:
	$(call print_py_title,Day 2 Part 2)
	@cd 2025/python && python main.py d2p2

# Day 3
python-d3p1:
	$(call print_py_title,Day 3 Part 1)
	@cd 2025/python && python main.py d3p1

python-d3p2:
	$(call print_py_title,Day 3 Part 2)
	@cd 2025/python && python main.py d3p2

# Day 4
python-d4p1:
	$(call print_py_title,Day 4 Part 1)
	@cd 2025/python && python main.py d4p1

python-d4p2:
	$(call print_py_title,Day 4 Part 2)
	@cd 2025/python && python main.py d4p2

# Day 5
python-d5p1:
	$(call print_py_title,Day 5 Part 1)
	@cd 2025/python && python main.py d5p1

python-d5p2:
	$(call print_py_title,Day 5 Part 2)
	@cd 2025/python && python main.py d5p2

# All Days
python-all: python-d1p1 python-d1p2 python-d2p1 python-d2p2 python-d3p1 python-d3p2 python-d4p1 python-d4p2 python-d5p1 python-d5p2

# ------------- GO -------------
# Day 1
go-d1p1:
	$(call print_go_title,Day 1 Part 1)
	@cd 2025/go && go run main.go d1p1

go-d1p2:
	$(call print_go_title,Day 1 Part 2)
	@cd 2025/go && go run main.go d1p2

# Day 2
go-d2p1:
	$(call print_go_title,Day 2 Part 1)
	@cd 2025/go && go run main.go d2p1

go-d2p2:
	$(call print_go_title,Day 2 Part 2)
	@cd 2025/go && go run main.go d2p2

# Day 3
go-d3p1:
	$(call print_go_title,Day 3 Part 1)
	@cd 2025/go && go run main.go d3p1

go-d3p2:
	$(call print_go_title,Day 3 Part 2)
	@cd 2025/go && go run main.go d3p2

# Day 4
go-d4p1:
	$(call print_go_title,Day 4 Part 1)
	@cd 2025/go && go run main.go d4p1

go-d4p2:
	$(call print_go_title,Day 4 Part 2)
	@cd 2025/go && go run main.go d4p2

# Day 5
go-d5p1:
	$(call print_go_title,Day 5 Part 1)
	@cd 2025/go && go run main.go d5p1

go-d5p2:
	$(call print_go_title,Day 5 Part 2)
	@cd 2025/go && go run main.go d5p2

# All Days
go-all: go-d1p1 go-d1p2 go-d2p1 go-d2p2 go-d3p1 go-d3p2 go-d4p1 go-d4p2 go-d5p1 go-d5p2

# ------------- TYPESCRIPT -------------
# Day 1
ts-d1p1:
	$(call print_ts_title,Day 1 Part 1)
	@cd 2025/ts && pnpm start d1p1

ts-d1p2:
	$(call print_ts_title,Day 1 Part 2)
	@cd 2025/ts && pnpm start d1p2

# Day 2
ts-d2p1:
	$(call print_ts_title,Day 2 Part 1)
	@cd 2025/ts && pnpm start d2p1

ts-d2p2:
	$(call print_ts_title,Day 2 Part 2)
	@cd 2025/ts && pnpm start d2p2

# Day 3
ts-d3p1:
	$(call print_ts_title,Day 3 Part 1)
	@cd 2025/ts && pnpm start d3p1

ts-d3p2:
	$(call print_ts_title,Day 3 Part 2)
	@cd 2025/ts && pnpm start d3p2

# Day 4
ts-d4p1:
	$(call print_ts_title,Day 4 Part 1)
	@cd 2025/ts && pnpm start d4p1

ts-d4p2:
	$(call print_ts_title,Day 4 Part 2)
	@cd 2025/ts && pnpm start d4p2

# Day 5
ts-d5p1:
	$(call print_ts_title,Day 5 Part 1)
	@cd 2025/ts && pnpm start d5p1

ts-d5p2:
	$(call print_ts_title,Day 5 Part 2)
	@cd 2025/ts && pnpm start d5p2

# All Days
ts-all: ts-d1p1 ts-d1p2 ts-d2p1 ts-d2p2 ts-d3p1 ts-d3p2 ts-d4p1 ts-d4p2 ts-d5p1 ts-d5p2

# ------------- RENDER -------------
render-render:
	@cd 2025/render && go run main.go render

render-init:
	@cd 2025/render && go run main.go init

# ------------- ALL LANGUAGES -------------
# Day 1
all-d1p1: render-init go-d1p1 python-d1p1 ts-d1p1 render-render
all-d1p2: render-init go-d1p2 python-d1p2 ts-d1p2 render-render
all-d1: render-init go-d1p1 go-d1p2 python-d1p1 python-d1p2 ts-d1p1 ts-d1p2 render-render

# Day 2
all-d2p1: render-init go-d2p1 python-d2p1 ts-d2p1 render-render
all-d2p2: render-init go-d2p2 python-d2p2 ts-d2p2 render-render
all-d2: render-init go-d2p1 go-d2p2 python-d2p1 python-d2p2 ts-d2p1 ts-d2p2 render-render

# Day 3
all-d3p1: render-init go-d3p1 python-d3p1 ts-d3p1 render-render
all-d3p2: render-init go-d3p2 python-d3p2 ts-d3p2 render-render
all-d3: render-init go-d3p1 go-d3p2 python-d3p1 python-d3p2 ts-d3p1 ts-d3p2 render-render

# Day 4
all-d4p1: render-init go-d4p1 python-d4p1 ts-d4p1 render-render
all-d4p2: render-init go-d4p2 python-d4p2 ts-d4p2 render-render
all-d4: render-init go-d4p1 go-d4p2 python-d4p1 python-d4p2 ts-d4p1 ts-d4p2 render-render

# Day 5
all-d5p1: render-init go-d5p1 python-d5p1 ts-d5p1 render-render
all-d5p2: render-init go-d5p2 python-d5p2 ts-d5p2 render-render
all-d5: render-init go-d5p1 go-d5p2 python-d5p1 python-d5p2 ts-d5p1 ts-d5p2 render-render

# All Days
all-all: render-init go-all python-all ts-all render-render
