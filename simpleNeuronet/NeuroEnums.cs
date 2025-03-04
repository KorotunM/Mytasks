using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _35_1_Korotun_NeurMih.NeiroMix
{
    enum NeuronType //тип нейрона
    {
        Hidden,  // нейон скрытого слоя
        Output //нейрон выходного слоя
    }

    enum NeuroworkMode //режим работы сети
    {
        Train,  // режим обучения
        Test, //режим тестирования
        Recogn //режим распознавания
    }

    enum MemoryMode //режим работы памяти
    {
        GET,  // считывание
        SET, // сохранение
        INIT // инициализация
    }
}
