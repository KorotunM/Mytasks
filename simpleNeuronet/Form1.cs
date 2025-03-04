using System;
using System.IO;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using _35_1_Korotun_NeurMih.NeiroMix;
using System.Windows.Forms.DataVisualization.Charting;

namespace _35_1_Korotun_NeurMih
{
    public partial class Form1 : Form
    {
        private double[] input_pixels = new double[15] { 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d, 0d };
    private Network network = new Network();
        private bool[] Party_time = new bool[15];
        private Timer timer;
        private int buttonClickCount = 0;
        private int totalButtons = 30;
        private Random random = new Random();
        public Form1()
        {
            InitializeComponent();

        }

        private void button1_Click(object sender, EventArgs e)
        {
            change_status(button1, 0);
        }
        private void change_status(Button b, int index)
        {
            if (b.BackColor == Color.Black)
            {
                b.BackColor = Color.White;
                input_pixels[index] = 0;
            }
            else
            {
                b.BackColor = Color.Black;
                input_pixels[index] = 1;
            }
            //label1.Text = index.ToString();    
        }

        private void button2_Click(object sender, EventArgs e)
        {
            change_status(button2, 1);
        }

        private void button3_Click(object sender, EventArgs e)
        {
            change_status(button3, 2);
        }

        private void button4_Click(object sender, EventArgs e)
        {
            change_status(button4, 3);
        }

        private void button5_Click(object sender, EventArgs e)
        {
            change_status(button5, 4);
        }

        private void button6_Click(object sender, EventArgs e)
        {
            change_status(button6, 5);
        }

        private void button7_Click(object sender, EventArgs e)
        {
            change_status(button7, 6);
        }

        private void button8_Click(object sender, EventArgs e)
        {
            change_status(button8, 7);
        }

        private void button9_Click(object sender, EventArgs e)
        {
            change_status(button9, 8);
        }
        private void button10_Click(object sender, EventArgs e)
        {
            change_status(button10, 9);
        }
        private void button11_Click(object sender, EventArgs e)
        {
            change_status(button11, 10);
        }
        private void button12_Click(object sender, EventArgs e)
        {
            change_status(button12, 11);
        }
        private void button13_Click(object sender, EventArgs e)
        {
            change_status(button13, 12);
        }
        private void button14_Click(object sender, EventArgs e)
        {
            change_status(button14, 13);
        }
        private void button15_Click(object sender, EventArgs e)
        {
            change_status(button15, 14);
        }

        private void Button_Click(object sender, EventArgs e)
        {
            buttonClickCount++;

            if (buttonClickCount == totalButtons)
            {
                timer.Start();
                stopButton.Visible = true;
            }

        }

        private void Timer_Tick(object sender, EventArgs e)
        {
            this.BackColor = Color.FromArgb(random.Next(256), random.Next(256), random.Next(256));
        }

        private void stopButton_Click(object sender, EventArgs e)
        {
            timer.Stop();
            stopButton.Visible = false;
        }
        //метод сохранения тестового примера
        private void SaveTest_Click(object sender, EventArgs e)
        {
            SaveTest_(numericUpDownExam.Value, input_pixels);
        }
        //метод сохранения обучающего примера
        private void SaveTrain_Click(object sender, EventArgs e)
        {
            SaveTrain_(numericUpDownExam.Value, input_pixels);//сделать аналогично для Test
        }
        // 
        private void numericUpDown1_ValueChanged(object sender, EventArgs e)
        {

        }
        private void SaveTrain_(decimal value, double[] input)
        {
            string pathDir = AppDomain.CurrentDomain.BaseDirectory;
            string nameFileTrain = Path.Combine(pathDir, "train.csv");

            // Формируем строку для записи
            string tmpStr = value.ToString() + ";";
            for (int i = 0; i < input.Length; i++)
            {
                tmpStr += input[i].ToString("F3") + ";"; // Ограничиваем до 3 знаков после запятой
            }

            // Проверка существования файла и запись
            if (!File.Exists(nameFileTrain))
            {
                File.WriteAllText(nameFileTrain, tmpStr + Environment.NewLine); // Создаем новый файл
            }
            else
            {
                File.AppendAllText(nameFileTrain, tmpStr + Environment.NewLine); // Добавляем строку в существующий файл
            }
        }

        private void SaveTest_(decimal value, double[] input)
        {
            string pathDir = AppDomain.CurrentDomain.BaseDirectory;
            string nameFileTest = Path.Combine(pathDir, "test.csv");

            // Формируем строку для записи
            string tmpStr = value.ToString() + ";";
            for (int i = 0; i < input.Length; i++)
            {
                tmpStr += input[i].ToString("F3") + ";"; // Ограничиваем до 3 знаков после запятой
            }

            // Проверка существования файла и запись
            if (!File.Exists(nameFileTest))
            {
                File.WriteAllText(nameFileTest, tmpStr + Environment.NewLine); // Создаем новый файл
            }
            else
            {
                File.AppendAllText(nameFileTest, tmpStr + Environment.NewLine); // Добавляем строку в существующий файл
            }
        }

        private void button16_Click(object sender, EventArgs e)
        {
            network.ForwardPass(network, input_pixels);
            labelOutput.Text = network.Fact.ToList().IndexOf(network.Fact.Max()).ToString();
            labelProbability.Text = (100*network.Fact.Max()).ToString("0.00")+"%";
        }

        private void button17_Click(object sender, EventArgs e)
        {
            network.Train(network);
            for (int i = 0; i < network.E_error_avr.Length; i++)
            {
                chartEavr.Series[0].Points.AddY(network.E_error_avr[i]);
            }
            MessageBox.Show("Обучение успешно завершено.", "Информация",
                MessageBoxButtons.OK, MessageBoxIcon.Information);
        }

        private void buttonTest_Click(object sender, EventArgs e)
        {
            network.Test(network);
            for (int i = 0; i< network.E_error_avr.Length; i++)
            {
                chartEavr.Series[0].Points.AddY(network.E_error_avr[i]);
            }
            MessageBox.Show("Тестирование успешно завершено.", "Информация",
                MessageBoxButtons.OK, MessageBoxIcon.Information);
        }


    }
}