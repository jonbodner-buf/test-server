update_deps:
	../buf/buf dep update
	../buf/buf generate --include-imports

update_deps_prod:
	BUF_USE_GIT_BRANCH_AS_LABEL=off ../buf/buf dep update
	../buf/buf generate --include-imports