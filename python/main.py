import sys
import importlib
import datetime
# import traceback
import os

if len(sys.argv) < 2:
    print("Usage:{} num-of-leetcode".format(sys.argv[0]))

# mod = importlib.import_module(sys.argv[1])
try:
    mod = importlib.import_module(sys.argv[1])
except ModuleNotFoundError:
    print("{}.py is not found".format(sys.argv[1]))
    sys.exit(1)
except:
    print("load {}.py fail".format(sys.argv[1]))
    sys.exit(1)

solution = None
try:
    solution = mod.Solution()
except AttributeError:
    print("{}.py is not implement Solution class".format(sys.argv[1]))
    sys.exit(1)
except Exception as e:
    print("Init solution fail: {}".format(str(e)))

try:
    start = datetime.datetime.now().timestamp()
    solution.run()
    end = datetime.datetime.now().timestamp()
    print("Use {0:.5f} ms".format((end - start)*1000))
except AttributeError:
    print("{}.py is not implement run()".format(sys.argv[1]))
    sys.exit(1)
except Exception as e:
    exc_type, exc_obj, exc_tb = sys.exc_info()
    filename = os.path.basename(exc_tb.tb_next.tb_frame.f_code.co_filename)
    template = "Run() has exception of type {0}. Arguments: {1!r} at {2}:{3}"
    message = template.format(type(e).__name__, e.args,
                              filename, exc_tb.tb_lineno)
    print(message)
