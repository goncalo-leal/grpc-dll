using System;
using System.Runtime.InteropServices;

class Program
{
    // Define the callback signature from the Go DLL
    [UnmanagedFunctionPointer(CallingConvention.Cdecl)]
    public delegate void Callback(string data);
    
    // Import the Go DLL functions
    [DllImport("lib/fixture-bridge.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern void SetCallback(Callback callback);

    [DllImport("lib/fixture-bridge.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern void StartServer(string port);


    // The callback function that will process data
    public static void ProcessData(string data)
    {
        Console.WriteLine($"Received data: {data}");
    }

    static void Main(string[] args)
    {
        Console.WriteLine("C# application is starting...");
        // Register the callback function
        SetCallback(ProcessData);

        Console.WriteLine("C# application is ready to receive data...");
        // Start the gRPC server in the DLL
        StartServer(":8080");

        Console.ReadLine();
    }
}
