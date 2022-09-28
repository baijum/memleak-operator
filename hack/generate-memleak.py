import os
from string import Template
import tempfile
import subprocess
import argparse

memleak_template = """\
apiVersion: cache.example.com/v1
kind: Memleak
metadata:
  name: mem-leak-${count}
spec:
  size: 3
"""

def apply_config(tmpl, **values):
    with tempfile.TemporaryDirectory(prefix="memleak-create-") as tmpdir:
        content = Template(tmpl).substitute(values)
        config_path = os.path.join(tmpdir, "config.yaml")
        with open(config_path, "w") as fd:
            fd.write(content)
        out = subprocess.run(["oc", "apply", "-f", config_path], capture_output=True)
        stdout = out.stdout.decode("utf-8")
        if out.returncode != 0:
            stderr = out.stderr.decode("utf-8")
        else:
            stderr = ""

    return stdout, stderr

def delete_config(tmpl, **values):
    with tempfile.TemporaryDirectory(prefix="memleak-delete-") as tmpdir:
        content = Template(tmpl).substitute(values)
        config_path = os.path.join(tmpdir, "config.yaml")
        with open(config_path, "w") as fd:
            fd.write(content)
        out = subprocess.run(["oc", "delete", "-f", config_path], capture_output=True)
        stdout = out.stdout.decode("utf-8")
        if out.returncode != 0:
            stderr = out.stderr.decode("utf-8")
        else:
            stderr = ""

    return stdout, stderr

def main():
    number = 500
    parser = argparse.ArgumentParser()
    parser.add_argument("-c", "--create", dest="create", type=str, required=False,
                                        help="create memleak services")
    parser.add_argument("-d", "--delete", dest="delete", type=str, required=False,
                                        help="delete memleak service")
    args = parser.parse_args()

    if args.create:
        for i in range(1, number):
            stdout, stderr = apply_config(memleak_template, count=str(i))
            print("stdout:\n", stdout, sep="")
            if stderr.strip():
                print("[ERROR] ", stderr)
    elif args.delete:
        for i in range(1, number):
            stdout, stderr = delete_config(memleak_template, count=str(i))
            print("stdout:\n", stdout, sep="")
            if stderr.strip():
                print("[ERROR] ", stderr)
    else:
        parser.print_help()

if __name__ == "__main__":
    main()
