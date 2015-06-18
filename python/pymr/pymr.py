import os
import pickle
import fnmatch
from subprocess import call

import pathlib
import click


@click.group()
def pymr():
    pass


@pymr.command()
@click.option('--directory', '-d', default='./')
@click.option('--tag', '-t', multiple=True)
@click.option('--append', is_flag=True)
def register(directory, tag, append):
    '''
    register a directory
    '''

    pymr_file = os.path.join(directory, '.pymr')
    new_tags = tag

    if append:
        if os.path.exists(pymr_file):
            cur_tags = pickle.load(open(pymr_file))
            new_tags = tuple(set(new_tags + cur_tags))

    pickle.dump(new_tags, open(pymr_file, 'wb'))


@pymr.command()
@click.argument('command')
@click.option('--basepath', '-b', default='./')
@click.option('--tag', '-t')
@click.option('--dryrun', is_flag=True)
def run(command, basepath, tag, dryrun):
    '''
    run a given command in matching sub-directories
    '''

    for fn in pathlib.Path(basepath).glob('**/.pymr'):

        with fn.open(mode='rb') as f:
            cur_tags = pickle.load(f)

        parent_dir = str(fn.parent)

        if tag in cur_tags:
            if dryrun:
                print('Would run {0} in {1}'.format(command, parent_dir))
            else:
                os.chdir(parent_dir)
                call(command, shell=True)


if __name__ == '__main__':
    pymr()
