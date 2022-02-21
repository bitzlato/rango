# frozen_string_literal: true

set :stage, 's6'
set :application, -> { 'rango-' + fetch(:stage).to_s }
set :deploy_to, -> { "/home/#{fetch(:user)}/#{fetch(:stage)}/#{fetch(:application)}" }

server ENV.fetch( 'STAGING_SERVER_6' ), user: fetch(:user), roles: fetch(:roles)