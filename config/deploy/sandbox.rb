# frozen_string_literal: true

set :stage, 'sandbox'
set :application, -> { 'rango' }
set :deploy_to, -> { "/home/#{fetch(:user)}/#{fetch(:stage)}/#{fetch(:application)}" }

server ENV.fetch( 'SANDBOX_SERVER' ), user: fetch(:user), roles: fetch(:roles)
