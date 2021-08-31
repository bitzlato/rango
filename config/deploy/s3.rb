# frozen_string_literal: true

set :build_domain, 'market-s3.bitzlato.com'
set :stage, 's3'

server ENV.fetch( 'STAGING_SERVER_2' ), user: fetch(:user), roles: fetch(:roles)
