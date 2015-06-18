require 'find'
require 'yaml'
require 'thor'

# [Ru]by [m]ulti-[r]epository command line tool
class Rumr < Thor
  desc 'register', 'Register a directory'
  method_option :directory,
                aliases: '-d',
                type: :string,
                default: './',
                desc: 'Directory to register'
  method_option :tag,
                aliases: '-t',
                type: :array,
                default: 'default',
                desc: 'Tag/s to register'
  method_option :append,
                type: :boolean,
                desc: 'Append given tags to any existing tags?'
  def register
    new_tags = options[:tag]
    rumr_file = File.join(options[:directory], '.rumr')
    if options[:append] && File.exist?(rumr_file)
      new_tags |= YAML.load_file(rumr_file)
    end
    IO.write(rumr_file, new_tags.to_yaml)
  end

  desc 'exec COMMAND', 'Execute (run) a given command on registered directories'
  method_option :basepath,
                aliases: '-b',
                type: :string,
                default: './',
                desc: 'Directory to begin search for rumr files.'
  method_option :tag,
                aliases: '-t',
                type: :string,
                default: 'default',
                desc: 'Tag to match against'
  method_option :dryrun,
                type: :boolean,
                desc: 'Display (do not execute) the commands to run.'
  def exec(command)
    Dir[File.join(options[:basepath], '**/.rumr')].each do |path|
      next unless YAML.load_file(path).include? options[:tag]
      if options['dryrun']
        puts "Would execute #{command} in #{path}"
      else
        Dir.chdir(File.dirname(path)) { puts `#{command}` }
      end
    end
  end
end
