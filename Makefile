.PHONY: gh\:setenv gh\:list gh\:delete

gh\:setenv:
	gh secret set -f .env.gh_secrets

gh\:list:
	gh secret list

gh\:delete:
	gh secret delete $(NAME)
