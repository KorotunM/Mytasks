using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;
using static System.Math;

namespace _35_1_Korotun_NeurMih.NeiroMix
{
    internal class Neuron
    {
        // поля с маленькой буквы, свойства с большой, поля приват, свойства паблик

        private NeuronType _type; // тип нейрона
        private double[] _weights; // на нулевом порог, дальше синоптические веса
        private double[] _inputs; // входные данные, на 1 элемент меньше чем в весах
        private double _output; // выход
        private double _derivative; // производная функции активации
        private double a = 0.01; //параметры для функции активации
        // функция активации геперболический тангенс

        public double[] Weights { get => _weights; set => _weights = value; }
        public double[] Inputs { get => _inputs; set => _inputs = value; }
        public double Output { get => _output; }
        public double Derivative { get => _derivative; }

        // Конструктор
        public Neuron(double[] weights, NeuronType type)
        {
            _type = type;
            _weights = weights;
        }

        // Метод активации нейрона (нелинейное преобразование водного слоя)
        public void Activator(double[] i)
        {
            _inputs = i; // передача вектора входного сигнала в массив входных данных
            double sum = _weights[0]; //аффинное преобразование чеез смещение (нулевой вес, порог)
            for (int j = 0; j < _inputs.Length; j++)
                sum += _inputs[j] * _weights[j + 1]; //линейные преобразования
            switch (_type)
            {
                case NeuronType.Hidden: // для нейронов скрытого слоя
                    _output = LeakyReLU(sum);
                    _derivative = LeakyReLU_Derivativator(sum);
                    break;
                case NeuronType.Output:
                    _output = Exp(sum);
                    break;
            }
        }

        // Функция активации Leaky ReLU
        public double LeakyReLU(double sum)
        {
            return (sum >= 0) ? sum : a * sum;
        }

        // Производная Leaky ReLU
        public double LeakyReLU_Derivativator(double sum)
        {
            return (sum >= 0) ? 1 : a;
        }
        //входной слой(15 входных нейронов т.к. поле 3х5) + скрытый слой(76) + скрытый слой(30) + выходной слой(10 выходных нейронов т.к 10 цифр)
        //Марина : 15 + 74 + 31 + 10

    }
}
