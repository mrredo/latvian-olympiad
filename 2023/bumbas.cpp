#include <iostream>
#include <cmath>

struct Pos {
    int x;
    int y;
    std::string toString() {
        return std::to_string(x) + " " + std::to_string(y);
    }
};

int maxLayers(int blocks) {
    return static_cast<int>(std::sqrt(static_cast<double>(blocks)));
}

int ballsUsed(int layers) {
    return (layers * (1 + (2 * layers - 1))) / 2;
}

Pos LastPosOfBall(int layers, int ballsleft) {
    if (ballsleft == 0) {
        return {0,layers};
    }
    if (layers == ballsleft) {
        Pos p = {layers - ballsleft + 1};
        p.y = layers - p.x + 1;
        return p;
    } else if (layers > ballsleft) {
        Pos p = {layers - ballsleft + 1};
        p.y = layers - p.x + 1;
        return p;
    } else if (ballsleft > layers) {
        Pos p = { -2 * layers + ballsleft - 1};
        p.y = layers - (p.x * -1) + 1;
        return p;
    }
    return {};
}

int main() {
    int blocktests[] = {9};
    int size = sizeof(blocktests) / sizeof(blocktests[0]);

    for (int i = 0; i < size; i++) {
        int v = blocktests[i];
        int l = maxLayers(v);
        int tb = ballsUsed(l);
        int bl = v - tb;
        std::cout << LastPosOfBall(l, bl).toString() << std::endl;
    }
    return 0;
}
