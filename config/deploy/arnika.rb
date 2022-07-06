# frozen_string_literal: true

set :stage, :production

server ENV['PRODUCTION_SERVER_ARNIKA'],
       user: fetch(:user),
       roles: fetch(:roles),
       ssh_options: { forward_agent: true }
