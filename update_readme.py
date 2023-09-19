import os
import glob
from jinja2 import Environment, FileSystemLoader

files = glob.glob('services_*.txt')
suffix = os.getenv('SUFFIX', '')

files_str = '\n'.join([suffix + file for file in files])

file_loader = FileSystemLoader('./templates')
env = Environment(loader=file_loader)

template = env.get_template('README.md.j2')

output = template.render(content=files_str)

with open('./README.md', 'w') as f:
    f.write(output)
