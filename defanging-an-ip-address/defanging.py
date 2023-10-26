import numpy as np

def defanging(address) -> str:
    convertedIp = convertIP(address)
    return convertedIp

def convertIP(address) -> str:
    period = "."
    leftBraces = "["
    rightBraces = "]"
    # strArr = set()
    strArr = np.array([])
    
    for addr in address:
       
        if addr == period:
            strArr.append(leftBraces)
            continue
        print(strArr)
        
    print(strArr)
    return ''.join(strArr)

if __name__ == '__main__':
    res1 = defanging("1.1.1.1")
    print(res1)