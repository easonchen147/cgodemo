#include "cgodemo.h"

int demo(int *data) {
    unsigned long count = 11;
    unsigned long nId0 = 1, nId1 = 1, nId2 = 1;
    int nLevel0 = 1, nLevel1 = 1, nLevel2 = 1;
    int i = 1;
    int offset = 14;
    for(i=1; i<10; i++) {
        data[0] = i;
        data[2] = count * nId0;
        data[4] = i * nLevel0;
        data[6] = count * nId1;
        data[8] = i * nLevel1;
        data[10] = count * nId2;
        data[12] = i * nLevel2;
        data += offset;
        count++;
    }
    return 0;
}

int demoLong(long *data) {
    unsigned long count = 21;
    unsigned long nId0 = 1, nId1 = 1, nId2 = 1;
    int nLevel0 = 1, nLevel1 = 1, nLevel2 = 1;
    int i = 1;
    int offset = 7;
    for(i=1; i<10; i++) {
        data[0] = i;
        data[1] = count * nId0;
        data[2] = i * nLevel0;
        data[3] = count * nId1;
        data[4] = i * nLevel1;
        data[5] = count * nId2;
        data[6] = i * nLevel2;
        data += offset;
        count++;
    }
    return 0;
}