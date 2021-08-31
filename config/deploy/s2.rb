# frozen_string_literal: true

set :stage, 's2'
set :application, -> { 'rango-' + fetch(:stage).to_s }

server ENV.fetch( 'STAGING_SERVER_2' ), user: fetch(:user), roles: fetch(:roles)
