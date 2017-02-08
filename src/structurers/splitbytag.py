import re
import sys, getopt
from bs4 import BeautifulSoup

def main(argv):

    # Default parameters
    input_location = ''
    output_path = ''

    try:
        opts, args = getopt.getopt(argv, "hi:o:", ["help", "input=", "output="])
    except getopt.GetoptError:
        usage()
        sys.exit(2)
    for opt, arg in opts:
        if opt in ("-h", "--help"):
            usage()
            sys.exit()
        elif opt in ("-i", "--input"):
            input_location = arg
        elif opt in ("-o", "--output"):
            if arg[-1] is not '/':
                output_path = arg + '/'
            else:
                output_path = arg
    if not opts:
        usage()
        sys.exit(2)

    if output_path == '' or input_location == '':
        print("input and output are mandatory parameters.")
        usage()
        sys.exit(2)

    return input_location, output_path


def usage():
    print("usage: splitbytags [--help] [--input=<path and file>] [--output=<path>]")
    print("--help: prints this usage guide and exits.")
    print("--input=<path and file>: pass in the path and file name of the input html.")
    print("--output=<path>: pass in the path where outputs are to be stored.")


def validate(htmltext):
    print(htmltext)
    tags_pattern = re.compile(r'(<\w+)')
    for tag in re.findall(tags_pattern, htmltext):
        return tag

def split(htmlfile):
    soup = BeautifulSoup(htmlfile, 'html5lib')
    return soup


if __name__ == "__main__":
    input_location, output_location = main(sys.argv[1:])
    htmlfile = open(input_location, 'r')
    htmltext = htmlfile.read().replace('\n', '')
    print(validate(htmltext))
