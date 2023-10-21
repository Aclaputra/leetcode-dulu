import unittest
from collections import defaultdict
import time

def add_two_integers(n1, n2):    
    return n1 + n2

class Test(unittest.TestCase):
   
    test_cases = [
        (10, 5, 15),
        (1000, 5000, 6000),
        (1999, 1, 2000)
    ]
    test_functions = [
        add_two_integers,
    ]
    
    def test_two_integers(self):
        num_runs = 1000
        function_runtimes = defaultdict(float)
        
        for n1, n2, expected in self.test_cases:
            for add_two_integers in self.test_functions:
                start = time.perf_counter()
                assert (
                    add_two_integers(n1, n2) == expected
                ), f"{add_two_integers.__name__} failed for value: {n1} {n2}"
                function_runtimes[add_two_integers.__name__] += (
                    time.perf_counter() - start
                ) * 1000
                
        print(f"\n{num_runs} runs")
        for function_name, runtime in function_runtimes.items():
            print(f"{function_name}: {runtime:.1f}ms")
                

if __name__ == '__main__':
    unittest.main()
    