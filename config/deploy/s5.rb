# frozen_string_literal: true

set :public_url, "https://#{ENV.fetch( 'STAGING_SERVER_5' )}/"
set :build_domain, ENV.fetch( 'STAGING_SERVER_5' )
set :stage, 's5'
set :application, -> { 'rango-' + fetch(:stage).to_s }

server ENV.fetch( 'STAGING_SERVER_5' ), user: fetch(:user), roles: fetch(:roles)