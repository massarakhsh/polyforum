using System;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace PolyWeb {
	public class PolyWeb {
		protected string Server = null;
		protected string User = null;
		protected string Pass = null;

		protected PolyWeb(string server, string user, string pass) {
			Server = server;
			User = user;
			Pass = pass;
		}

		static public PolyWeb Build(string server, string user, string pass) {
			PolyWeb poly = new PolyWeb(server, user, pass);
			return poly;
		}

		public string Get(string Table, int Id) {
			return sendRequest("get", Table, Id, null);
		}

		public string GetListAll(string Table) {
			return sendRequest("list", Table, 0, null);
		}

		public string Update(string Table, int Id, string data) {
			return sendRequest("update", Table, Id, data);
		}

		public string Insert(string Table, string data) {
			return sendRequest("insert", Table, 0, data);
		}

		public string LastId(string Table) {
			return sendRequest("lastid", Table, 0, null);
		}

		public string Delete(string Table, int Id) {
			return sendRequest("delete", Table, Id, null);
		}

		public string Ean13(string code) {
			return sendRequest("ean13", null, 0, "data=" + code);
		}

		public string QRCode(string Table, string code) {
			return sendRequest("qrcode", Table, 0, "data=" + code);
		}

		public string Search(string Table, string data) {
			return Search(Table, data, null);
		}
		public string Search(string Table, string data, string sort) {
			string parm = "";
			if (data != null) {
				if (parm.Length > 0) parm += "&";
				parm += "data=" + Base64Encode(data);
			}
			if (sort != null) {
				if (parm.Length > 0) parm += "&";
				parm += "sort=" + Base64Encode(sort);
			}
			return sendRequest("search", Table, 0, parm);
		}

		public string SearchForum(int IdForum, string Table) {
			return SearchForum(IdForum, Table, null, null, null);
		}
		public string SearchForum(int IdForum, string Table, string data) {
			return SearchForum(IdForum, Table, null, data, null);
		}
		public string SearchForum(int IdForum, string Table, string data, string sort) {
			return SearchForum(IdForum, Table, null, data, sort);
		}
		public string SearchForum(int IdForum, string Table, string field, string data, string sort) {
			string parm = "";
			if (field != null) {
				if (parm.Length > 0) parm += "&";
				parm += "field=" + Base64Encode(field);
			}
			if (data != null) {
				if (parm.Length > 0) parm += "&";
				parm += "data=" + Base64Encode(data);
			}
			if (sort != null) {
				if (parm.Length > 0) parm += "&";
				parm += "sort=" + Base64Encode(sort);
			}
			return sendRequest("searchforum", Table, IdForum, parm);
		}

		public string SendSMSCode(string phonenumber) {
			return sendRequest("sendsmscode", onlyDigits(phonenumber), 0, null);
		}

		public string ProbeSMSCode(string phonenumber, string code) {
			return sendRequest("probesmscode", onlyDigits(phonenumber), 0, "data=" + code);
		}

		public string CreatePerson(string phonenumber) {
			return CreatePerson(phonenumber, null);
		}
		public string CreatePerson(string phonenumber, string data) {
			return sendRequest("createperson", onlyDigits(phonenumber), 0, data);
		}

		public string SeekPerson(string login, string password) {
			string parm = "";
			if (login != null) {
				if (parm.Length > 0) parm += "&";
				parm += "login=" + Base64Encode(login);
			}
			if (password != null) {
				if (parm.Length > 0) parm += "&";
				parm += "password=" + Base64Encode(password);
			}
			return sendRequest("seekperson", null, 0, parm);
		}

		protected string sendRequest(string cmd, string table, int id, string data) {
			string sql = Server + "/api";
			if (cmd != null && cmd.Length > 0) {
				sql += "/" + cmd;
			}
			if (table != null && table.Length > 0) {
				sql += "/" + table;
			}
			if (id > 0) {
				sql += "/" + id;
			}
			System.Net.WebRequest req = System.Net.WebRequest.Create(sql);
			if (User != null || Pass != null) {
				string pair = User + ":" + Pass;
				string auth = "Basic " + Convert.ToBase64String(Encoding.UTF8.GetBytes(pair));
				req.Headers.Add("Authorization", auth);
			}
			req.Method = "POST";
			req.Timeout = 15000;
			req.ContentType = "application/x-www-form-urlencoded";
			if (data != null && data.Length > 0) {
				//byte[] sentData = System.Text.Encoding.GetEncoding(1251).GetBytes(Data);
				byte[] sentData = Encoding.UTF8.GetBytes(data);
				req.ContentLength = sentData.Length;
				System.IO.Stream sendStream = req.GetRequestStream();
				sendStream.Write(sentData, 0, sentData.Length);
				sendStream.Close();
			}
			System.Net.WebResponse res = req.GetResponse();
			System.IO.Stream ReceiveStream = res.GetResponseStream();
			System.IO.StreamReader sr = new System.IO.StreamReader(ReceiveStream, Encoding.UTF8);
			Char[] read = new Char[1024];
			int count = sr.Read(read, 0, 1024);
			string Out = String.Empty;
			while (count > 0) {
				String str = new String(read, 0, count);
				Out += str;
				count = sr.Read(read, 0, 1024);
			}
			sr.Close();
			ReceiveStream.Close();
			res.Close();
			return Out;
		}

		protected string onlyDigits(string code) {
			Regex rgx = new Regex("\\D");
			return rgx.Replace(code, "");
		}

		public static string Base64Encode(string plainText) {
			var plainTextBytes = System.Text.Encoding.UTF8.GetBytes(plainText);
			var str = System.Convert.ToBase64String(plainTextBytes);
			return str.Replace("=", "@");
		}
	}
}
