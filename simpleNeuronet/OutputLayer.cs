﻿
using System;

namespace _35_1_Korotun_NeurMih.NeiroMix
{
    class OutputLayer : Layer
    {
        public OutputLayer(int non, int nopn, NeuronType nt, string type) : base(non, nopn, nt, type) { }

        public override void Recognize(Network net, Layer nextLayer)
        {
            double e_sum = 0;
            for (int i = 0; i < neurons.Length; i++)
                e_sum += neurons[i].Output;
            for (int i = 0; i < neurons.Length; i++)
                net.Fact[i] = neurons[i].Output / e_sum;
            // все значения должны быть от нуля до единицы, а сумма меньше 1
        }

        public override double[] BackwardPass(double[] errors)
        {
            double[] gr_sum = new double[numofprevneurons + 1];
            for (int j = 0; j < numofprevneurons + 1; j++)
            {
                double sum = 0;
                for (int k = 0; k < numofneurouns; k++)
                    sum += neurons[k].Weights[j] * errors[k];
                gr_sum[j] = sum;
            }
            for (int i = 0; i < numofneurouns; i++)
            {
                if (neurons[i] == null)
                    throw new InvalidOperationException($"Нейрон {i} не инициализирован.");

                if (neurons[i].Inputs == null)
                    throw new InvalidOperationException($"Inputs для нейрона {i} не инициализированы.");

                for (int n = 0; n < numofprevneurons + 1; n++)
                {
                    double deltaw;

                    if (lastdeltaweights == null || lastdeltaweights.GetLength(0) <= i || lastdeltaweights.GetLength(1) <= n)
                        throw new InvalidOperationException("Массив lastdeltaweights не корректно инициализирован.");

                    if (n == 0)
                    {
                        deltaw = momentum * lastdeltaweights[i, 0] + learnigrate * errors[i];
                    }
                    else
                    {
                        deltaw = momentum * lastdeltaweights[i, n] + learnigrate * neurons[i].Inputs[n - 1] * errors[i];
                    }

                    lastdeltaweights[i, n] = deltaw;
                    neurons[i].Weights[n] += deltaw;
                }
            }
            return gr_sum;
        }
    }
}
