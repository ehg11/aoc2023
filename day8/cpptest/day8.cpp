#include <cctype>
#include <cstdio>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>
using namespace std;

struct Destinations {
    Destinations() {
        this->left = "";
        this->right = "";
    }
    Destinations(string left, string right) {
        this->left = left;
        this->right = right;
    }
    string left;
    string right;
};

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

string trimSpace(const string &originalStr) {
    string trimmed = "";
    size_t firstNonSpace = 0;
    size_t lastNonSpace = originalStr.size() - 1;

    while (true) {
        if (originalStr[firstNonSpace] != ' ') {
            break;
        }
        firstNonSpace++;
    }
    while (true) {
        if (originalStr[lastNonSpace] != ' ') {
            break;
        }
        lastNonSpace--;
    }

    for (size_t i = firstNonSpace; i <= lastNonSpace; i++) {
        trimmed += originalStr[i]; 
    }

    return trimmed;
}

pair<string, string> getDestinations(const string &dests) {
    string left = "";
    string right = "";

    vector<string> destsSplit = split(dests, ',');
    for (size_t i = 0; i < destsSplit[0].size(); i++) {
        if (isalnum(destsSplit[0][i])) {
            left += destsSplit[0][i];
        }
    }
    for (size_t i = 0; i < destsSplit[1].size(); i++) {
        if (isalnum(destsSplit[1][i])) {
            right += destsSplit[1][i];
        }
    }

    return make_pair(left, right);
}

unordered_map<string, Destinations> parseLines(const vector<string> &translations) {
    unordered_map<string, Destinations> destMap;
    for (const string& bufline : translations) {
        vector<string> buflineSplit = split(bufline, '=');
        string key = trimSpace(buflineSplit[0]);
        string dests = trimSpace(buflineSplit[1]);
        pair<string, string> destPair = getDestinations(dests);
        Destinations destStruct = Destinations(destPair.first, destPair.second);
        destMap[key] = destStruct;
    }

    return destMap;
}

vector<string> getStartPositions(const unordered_map<string, Destinations> &destMap) {
    vector<string> startPositions;
    for (auto it = destMap.begin(); it != destMap.end(); it++) {
        string position = it->first;
        if (position[position.size() - 1] == 'A') {
            startPositions.push_back(position);
        }
    }

    return startPositions;
}

bool allOnZ(vector<string> &positions) {
    for (const string& pos : positions) {
        if (pos[pos.size() - 1] != 'Z') {
            return false;
        }
    }
    return true;
}

template <typename T>
void show(vector<T> v) {
    for (const T& t : v) {
        cout << t << " ";
    }
    printf("\n");
}

int getNumSteps(
    const string &steps,
    const unordered_map<string, Destinations> &destMap,
    const vector<string> &startPositions
) {
    int numSteps = 0;
    vector<string> positions(startPositions);

    while (!allOnZ(positions)) {
        int index = numSteps % steps.size();
        numSteps++;
        char direction = steps[index];
        switch (direction) {
            case 'L': {
                for (size_t i = 0; i < positions.size(); i++) {
                    string key = positions[i];
                    positions[i] = destMap.find(key)->second.left;
                }
                break;
            }
            case 'R': {
                for (size_t i = 0; i < positions.size(); i++) {
                    string key = positions[i];
                    positions[i] = destMap.find(key)->second.right;
                }
                break;
            }
            default: continue;
        }
        printf("%d ", numSteps);
        show(positions);
    }
    
    return numSteps;
}

int main() {
    ifstream inputfile("d8");
    string line;
    vector<string> buffer;
    while (getline(inputfile, line)) {
        buffer.push_back(line);
    }

    string steps = buffer[0];
    vector<string>::iterator sliceStart = buffer.begin() + 2;
    vector<string>::iterator sliceEnd = buffer.end();
    vector<string> translations(sliceStart, sliceEnd);

    unordered_map<string, Destinations> destMap = parseLines(translations);
    vector<string> startPositions = getStartPositions(destMap);

    int numSteps = getNumSteps(steps, destMap, startPositions);
    printf("%d\n", numSteps);
}
