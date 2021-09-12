# frozen_string_literal: true

set :stage, 'sandbox'
set :application, -> { 'rango-' + fetch(:stage).to_s }
set :deploy_to, -> { "/home/#{fetch(:user)}/#{fetch(:stage)}/#{fetch(:application)}" }

server ENV.fetch( 'STAGING_SERVER' ), user: fetch(:user), roles: fetch(:roles)
