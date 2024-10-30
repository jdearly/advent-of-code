#include <stdio.h>

int solve() {
	FILE *file = fopen("input.txt", "r");

	if (NULL == file) {
		printf("file can't be opened\n");
		return -1;
	}

	int floor = 0;
	char c;

	while ((c = fgetc(file)) != EOF) {
		if (c == ')') {
			floor--;
		}
		if (c == '(') {
			floor++;
		}
	}

	fclose(file);

	return floor;
}

int main()
{
	printf("Answer: %d\n", solve());
	return 0;
}
