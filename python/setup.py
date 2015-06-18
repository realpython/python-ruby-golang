from setuptools import setup, find_packages

classifiers = [
    'Environment :: Console',
    'Operating System :: OS Independent',
    'License :: OSI Approved :: GNU General Public License v3 or later (GPLv3+)',
    'Intended Audience :: Developers',
    'Programming Language :: Python',
    'Programming Language :: Python :: 2',
    'Programming Language :: Python :: 2.7'
]

setuptools_kwargs = {
    'install_requires': [
        'click>=4,<5',
        'pathlib>=1,<2'
    ],
    'entry_points': {
        'console_scripts': [
            'pymr = pymr.pymr:pymr',
        ]
    }
}


setup(
    name='pymr',
    description='A tool for executing ANY command in a set of tagged directories.',
    author='Kyle W Purdon',
    author_email='kylepurdon@gmail.com',
    url='https://github.com/kpurdon/pymr',
    download_url='https://github.com/kpurdon/pymr',
    version='2.0.1',
    packages=find_packages(),
    classifiers=classifiers,
    **setuptools_kwargs
)
