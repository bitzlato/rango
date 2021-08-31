# frozen_string_literal: true

set :stage, :staging
set :build_domain, 'ex-stage.bitzlato.bz'

server ENV.fetch( 'STAGING_SERVER' ), user: fetch(:user), roles: fetch(:roles)
