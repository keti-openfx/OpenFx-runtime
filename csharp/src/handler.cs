using System.Text;
using System;

namespace Fx
{
    class Function
    {
        public byte[] Handler(byte[] Input)
        {
            string runtime = "[CSharp] ";
	    byte[] strbyte = Encoding.UTF8.GetBytes(runtime);

	    byte[] res = new byte[strbyte.Length+Input.Length];
	    Array.Copy(strbyte, 0, res, 0, strbyte.Length);
	    Array.Copy(Input, 0, res, strbyte.Length, Input.Length);

	    return res; 
        }
    }
}
