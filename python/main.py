import sys
import importlib

if len(sys.argv) < 2:
  print("Usage:{} num-of-leetcode".format(sys.argv[0]))

try:
  mod = importlib.import_module(sys.argv[1])
except:
  print("{}.py is not found".format(sys.argv[1]))
  sys.exit(1)

try:
  solution = mod.Solution()
  solution.run()
except:
  print("{}.py is not implement run()".format(sys.argv[1]))
  sys.exit(1)