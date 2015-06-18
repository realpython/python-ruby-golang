Gem::Specification.new do |s|
  s.name        = 'rumr'
  s.version     = '1.0.2'
  s.summary     = 'Run system commands in sets' \
                  ' of registered and tagged directories.'
  s.description = '[Ru]by [m]ulti-[r]epository Tool'
  s.authors     = ['Kyle W. Purdon']
  s.email       = 'kylepurdon@gmail.com'
  s.files       = ['lib/rumr.rb']
  s.homepage    = 'https://github.com/kpurdon/rumr'
  s.license     = 'GPLv3'
  s.executables << 'rumr'
  s.add_dependency('thor', ['~>0.19.1'])
end
