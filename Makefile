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

# All Days
python-all: python-d1p1 python-d1p2 python-d2p1 python-d2p2 python-d3p1 python-d3p2

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

# All Days
go-all: go-d1p1 go-d1p2 go-d2p1 go-d2p2 go-d3p1 go-d3p2

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

# All Days
ts-all: ts-d1p1 ts-d1p2 ts-d2p1 ts-d2p2 ts-d3p1 ts-d3p2

# ------------- ALL LANGUAGES -------------
# Day 1
all-d1p1: go-d1p1 python-d1p1 ts-d1p1
all-d1p2: go-d1p2 python-d1p2 ts-d1p2

# Day 2
all-d2p1: go-d2p1 python-d2p1 ts-d2p1
all-d2p2: go-d2p2 python-d2p2 ts-d2p2

# Day 3
all-d3p1: go-d3p1 python-d3p1 ts-d3p1
all-d3p2: go-d3p2 python-d3p2 ts-d3p2

# All Days
all-all: go-all python-all ts-all
