#include <iostream>
#include <iomanip>
#include <cmath>
#include <clocale>

using std::wcout; using std::wcin;
using std::round;

double CheckErr(double& val) {
	while (true) {
		if ((val > 20 || val < 1) && wcin.peek() != ' ' && !isalpha(wcin.peek())) {
			wcin.clear(); wcin.ignore(32767, '\n');
			wcout << L"Ошибка! Вне допустимого диапазона значений\n" <<
				L"Введите размерность матрицы (целое число от 0 до 20):\n";
			wcin >> val;
		}
		else if (wcin.fail() || isalpha(wcin.peek()) || (L'А' <= wchar_t(wcin.peek()) && L'я' >= wchar_t(wcin.peek()))) {
			wcin.clear(); wcin.ignore(32767, '\n');
			wcout << L"Ошибка! Введено не число\n" <<
				L"Введите размерность матрицы (целое число от 0 до 20):\n";
			wcin >> val;
		}
		else if (wcin.peek() == ' ' || wcin.peek() == ',' || wcin.peek() == '\t') {
			wcin.clear(); wcin.ignore(32767, '\n');
			wcout << L"Ошибка! Неккоректный ввод.\n" <<
				L"Введите размерность матрицы (целое число от 0 до 20):\n";
			wcin >> val;
		}
		else if (round(val) != val) {
			wcin.clear(); wcin.ignore(32767, '\n');
			wcout << L"Ошибка! Введено дробное число.\n" <<
				L"Введите размерность матрицы (целое число от 0 до 20):\n";
			wcin >> val;
		}
		else {
			return val;
		}
	}
}
int main() {
	srand(time(NULL));
	setlocale(LC_CTYPE, "rus");
	wcout << L"Лабораторная работа №3. Вариант №39\n" <<
		L"Сафонов А.П., группа 19-ИЭ-2\n" <<
		L"Задание: Модуль разности максимальных элементов главной и побочной диагоналей.\n\n";
	double check, max = -100, max2 = -100;
	int s;

	wcout << L"Введите размерность матрицы (целое число от 0 до 20):\n";
	wcin >> check; s = CheckErr(check);
	int v1 = s - 1, v = 0;

	double matrix[20][20];

	for (int i = 0; i < s; ++i) {
		for (int j = 0; j < s; ++j) {
			matrix[i][j] = rand() / double(201) - double(100);
		}
	}

	for (int j = 0; j < s; ++j) {
		for (int i = 0; i < s; ++i) {
			max = matrix[j][v] > max ? matrix[j][v] : max;
		}
		++v;
		wcout << '\n';
	}


	v1 = s - 1;
	for (int j = 0; j < s; ++j) {
		for (int i = 0; i < s; ++i) {
			wcout << std::setw(10) << matrix[j][i];
			max2 = matrix[j][v1] > max2 ? matrix[j][v1] : max2;

		}
		--v1;
		wcout << '\n';
	}

	wcout << '\n';
	wcout << L"Результат: ";
	double res = max - max2;
	if (res < 0)
		res = res * -1;
	wcout << res << '\n';
	wcout << max << '\n';
	wcout << max2 << '\n';

}