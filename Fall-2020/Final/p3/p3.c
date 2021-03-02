#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct count {
    char word[20];
    int cnt;
};

int search(struct count result[], int size, const char *word) {
    for (int i = 0; i < size; ++i) {
        if (strcmp(result[i].word, word) == 0)
            return i;
    }
    return -1;
}

void check_word(struct count result[], int *count, char *str, int len) {
    char word[20];
    strncpy(word, str, len);
    word[len] = '\0';
    int idx = search(result, *count, word);
    if (idx >= 0) {
        result[idx].cnt++;
    } else {
        strcpy(result[*count].word, word);
        result[(*count)++].cnt = 1;
    }
}

int frequency(int rows, int len, char strings[rows][len], struct count result[]) {
    int count = 0;
    for (int i = 0; i < rows; ++i) {
        char *index, *s = strings[i];
        while ((index = strstr(s, " ")) != NULL) {
            int l = index - s;
            if (l > 0) {
                check_word(result, &count, s, l);
            }
            s = index + 1;
        }
        if(strlen(s) > 0){
            check_word(result, &count, s, strlen(s));
        }
    }
    return count;
}

int main() {
    int rows, len;
    scanf("%d %d\n", &rows, &len);
    char lines[rows][len];
    for (int i = 0; i < rows; ++i) {
        gets(lines[i]);
    }
    int max_num_words = rows * len / 4;
    struct count result[max_num_words];
    int count = frequency(rows, len, lines, result);
    for (int i = 0; i < count; ++i) {
        printf("%s %d\n", result[i].word, result[i].cnt);
    }
    return 0;
}
