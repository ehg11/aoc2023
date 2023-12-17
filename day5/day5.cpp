#include <iostream>
#include <sstream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;
using u64 = unsigned long long;

struct Mapping {
    Mapping() {
        this->dst = 0;
        this->src = 0;
        this->length = 0;
    }
    Mapping(u64 dst, u64 src, u64 length) {
        this->dst = dst;
        this->src = src;
        this->length = length;
    }
    u64 dst;
    u64 src;
    u64 length;
};

template <typename T>
void show(vector<T> vec) {
    for (const T& item : vec) {
        cout << item << " ";
    }
    printf("\n");
}


vector<string> split(const string &s, char delimeter) {
    vector<string> result;
    stringstream ss(s);
    string item;

    while (getline(ss, item, delimeter)) {
        if (item.length() <= 0) {
            continue;
        } 
        result.push_back(item);
    }

    return result;
} 

vector<u64> initSeeds(string line) {
    string rhs = split(line, ':')[1];
    vector<string> seedStrs = split(rhs, ' ');

    vector<u64> seeds;
    u64 seedStart = 0;

    for (size_t i = 0; i < seedStrs.size(); i++) {
        string seedItem = seedStrs[i];
        u64 seedNum = stoull(seedItem);

        if (i % 2 == 0) {
             seedStart = seedNum;
        } else {
            for (u64 j = seedStart; j < seedStart + seedNum; j++) {
                seeds.push_back(j);
            }
        }
    }

    return seeds;
}

void initMap(vector<Mapping*> &mapping, string line) {
    vector<string> mappingStrs = split(line, ' ');
    vector<u64> mappingNums;

    for (const string& str : mappingStrs) {
        u64 num = stoull(str);
        mappingNums.push_back(num);
    }

    Mapping *newMapping = new Mapping(mappingNums[0], mappingNums[1], mappingNums[2]);
    mapping.push_back(newMapping);
}

u64 findCorrespondingValue(u64 key, vector<Mapping*> &mappings) {
    for (const Mapping* mapping : mappings) {
        u64 src = mapping->src;
        u64 end = src + mapping->length;

        if (key >= src && key < end) {
            u64 diff = key - src;
            return mapping->dst + diff;
        }
    }

    return key;
}

void clearMappings(vector<Mapping*>& currMappings) {
    for (const Mapping* mapping : currMappings) {
        delete mapping;
    }
    currMappings.clear();
}

int main() {
    ifstream inputFile("d5");
    string line;
    vector<string> buffer;
    while (getline(inputFile, line)) {
        buffer.push_back(line);
    }

    vector<u64> currKeys = initSeeds(buffer[0]);
    vector<Mapping*> currMappings;

    u64 lineNumber = 3;

    u64 translations = 7;
    while (translations > 0) {
        while (true) {
            line = buffer[lineNumber];
            if (line.size() == 0) {
                lineNumber += 2;
                break;
            }
            initMap(currMappings, line);
            lineNumber++;
        }

        for (size_t i = 0; i < currKeys.size(); i++) {
            currKeys[i] = findCorrespondingValue(currKeys[i], currMappings);
        }
        clearMappings(currMappings);
        printf("translations %llu\n", translations);
        translations--;
    }

    u64 minLocation = currKeys[0];
    for (size_t i = 1; i < currKeys.size(); i++) {
        if (currKeys[i] < minLocation) {
            minLocation = currKeys[i];
        }
    }
    printf("Min Location: %llu\n", minLocation);
}
