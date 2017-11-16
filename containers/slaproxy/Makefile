FILES := \
	../../auth_gen.go \
	../../auth_test.go \
	../../bots_gen.go \
	../../channels_gen.go \
	../../channels_test.go \
	../../chat_gen.go \
	../../chat_test.go \
	../../client.go \
	../../cmd/slaproxy/slaproxy.go \
	../../emoji_gen.go \
	../../emoji_test.go \
	../../events/event_types_gen.go \
	../../events/events.go \
	../../events/events_test.go \
	../../events/interface.go \
	../../groups_gen.go \
	../../groups_test.go \
	../../interface.go \
	../../internal/cmd/genevents/genevents.go \
	../../internal/cmd/genmethods/genmethods.go \
	../../internal/cmd/genscopes/genscopes.go \
	../../internal/cmd/genserver/genserver.go \
	../../logger.go \
	../../oauth_gen.go \
	../../oauth_test.go \
	../../objects/errors.go \
	../../objects/errors_test.go \
	../../objects/interface.go \
	../../objects/list.go \
	../../objects/timestamp.go \
	../../objects/user_profile.go \
	../../options.go \
	../../reactions_gen.go \
	../../reactions_test.go \
	../../reminders_gen.go \
	../../rtm/client.go \
	../../rtm/client_test.go \
	../../rtm/event_types.go \
	../../rtm/event_types_string.go \
	../../rtm/events.go \
	../../rtm/interface.go \
	../../rtm/option.go \
	../../rtm/rtm.go \
	../../rtm/rtm_example_test.go \
	../../rtm/rtm_test.go \
	../../rtm_gen.go \
	../../rtm_test.go \
	../../scopes_gen.go \
	../../sequence.go \
	../../sequence_test.go \
	../../server/mockserver/mockserver.go \
	../../server/mockserver/mockserver_test.go \
	../../server/mockserver/response.go \
	../../server/proxyserver/proxyserver.go \
	../../server/server.go \
	../../server/server_test.go \
	../../services.go \
	../../slack.go \
	../../slack_example_test.go \
	../../slack_test.go \
	../../usergroups_gen.go \
	../../usergroups_test.go \
	../../usergroups_users_gen.go \
	../../usergroups_users_test.go \
	../../users_gen.go \
	../../users_profile_gen.go \
	../../users_profile_test.go \
	../../users_test.go

slaproxy: $(FILES)
	@echo "building slaproxy binary (for linux/amd64)"
	@GOOS=linux GOARCH=amd64 go build ../../cmd/slaproxy/slaproxy.go

docker: slaproxy
	@echo "building docker image..."
	@# TODO: version it
	@docker build -t go-slack/slaproxy .