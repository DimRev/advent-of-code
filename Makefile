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
python-d1p1:
	$(call print_py_title,Day 1 Part 1)
	@cd 2025/python && python main.py d1p1

python-d1p2:
	$(call print_py_title,Day 1 Part 2)
	@cd 2025/python && python main.py d1p2

# ------------- GO -------------
go-d1p1:
	$(call print_go_title,Day 1 Part 1)
	@cd 2025/go && go run main.go d1p1

go-d1p2:
	$(call print_go_title,Day 1 Part 2)
	@cd 2025/go && go run main.go d1p2

# ------------- TYPESCRIPT -------------
ts-d1p1:
	$(call print_ts_title,Day 1 Part 1)
	@cd 2025/ts && pnpm start d1p1

ts-d1p2:
	$(call print_ts_title,Day 1 Part 2)
	@cd 2025/ts && pnpm start d1p2

# ALL LANGUAGES
all-d1p1: go-d1p1 python-d1p1 ts-d1p1
all-d1p2: go-d1p2 python-d1p2 ts-d1p2