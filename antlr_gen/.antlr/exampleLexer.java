// Generated from /Users/drbngd/Documents/UC/14_2024_SPR/EECE6083/DGFC/antlr_gen/example.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class exampleLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ADD_OP=1, MULT_OP=2, IF_KW=3, THEN_KW=4, ELSE_KW=5, WHILE_KW=6, DO_KW=7, 
		END_KW=8, COLON_EQUAL=9, L_BRAKET=10, R_BRAKET=11, L_PAREN=12, R_PAREN=13, 
		SEMI_COLON=14, ID=15, NUM=16, WS=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ADD_OP", "MULT_OP", "IF_KW", "THEN_KW", "ELSE_KW", "WHILE_KW", "DO_KW", 
			"END_KW", "COLON_EQUAL", "L_BRAKET", "R_BRAKET", "L_PAREN", "R_PAREN", 
			"SEMI_COLON", "ID", "NUM", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, "'if'", "'then'", "'else'", "'while'", "'do'", "'end'", 
			"':='", "'['", "']'", "'('", "')'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ADD_OP", "MULT_OP", "IF_KW", "THEN_KW", "ELSE_KW", "WHILE_KW", 
			"DO_KW", "END_KW", "COLON_EQUAL", "L_BRAKET", "R_BRAKET", "L_PAREN", 
			"R_PAREN", "SEMI_COLON", "ID", "NUM", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public exampleLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "example.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0011_\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\u000e\u0004\u000eP\b\u000e\u000b\u000e\f\u000eQ\u0001"+
		"\u000f\u0004\u000fU\b\u000f\u000b\u000f\f\u000fV\u0001\u0010\u0004\u0010"+
		"Z\b\u0010\u000b\u0010\f\u0010[\u0001\u0010\u0001\u0010\u0000\u0000\u0011"+
		"\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r"+
		"\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e"+
		"\u001d\u000f\u001f\u0010!\u0011\u0001\u0000\u0005\u0002\u0000++--\u0002"+
		"\u0000**//\u0001\u0000az\u0001\u000009\u0003\u0000\t\n\r\r  a\u0000\u0001"+
		"\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005"+
		"\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001"+
		"\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000"+
		"\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000"+
		"\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000"+
		"\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000"+
		"\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000"+
		"\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000"+
		"\u0000\u0001#\u0001\u0000\u0000\u0000\u0003%\u0001\u0000\u0000\u0000\u0005"+
		"\'\u0001\u0000\u0000\u0000\u0007*\u0001\u0000\u0000\u0000\t/\u0001\u0000"+
		"\u0000\u0000\u000b4\u0001\u0000\u0000\u0000\r:\u0001\u0000\u0000\u0000"+
		"\u000f=\u0001\u0000\u0000\u0000\u0011A\u0001\u0000\u0000\u0000\u0013D"+
		"\u0001\u0000\u0000\u0000\u0015F\u0001\u0000\u0000\u0000\u0017H\u0001\u0000"+
		"\u0000\u0000\u0019J\u0001\u0000\u0000\u0000\u001bL\u0001\u0000\u0000\u0000"+
		"\u001dO\u0001\u0000\u0000\u0000\u001fT\u0001\u0000\u0000\u0000!Y\u0001"+
		"\u0000\u0000\u0000#$\u0007\u0000\u0000\u0000$\u0002\u0001\u0000\u0000"+
		"\u0000%&\u0007\u0001\u0000\u0000&\u0004\u0001\u0000\u0000\u0000\'(\u0005"+
		"i\u0000\u0000()\u0005f\u0000\u0000)\u0006\u0001\u0000\u0000\u0000*+\u0005"+
		"t\u0000\u0000+,\u0005h\u0000\u0000,-\u0005e\u0000\u0000-.\u0005n\u0000"+
		"\u0000.\b\u0001\u0000\u0000\u0000/0\u0005e\u0000\u000001\u0005l\u0000"+
		"\u000012\u0005s\u0000\u000023\u0005e\u0000\u00003\n\u0001\u0000\u0000"+
		"\u000045\u0005w\u0000\u000056\u0005h\u0000\u000067\u0005i\u0000\u0000"+
		"78\u0005l\u0000\u000089\u0005e\u0000\u00009\f\u0001\u0000\u0000\u0000"+
		":;\u0005d\u0000\u0000;<\u0005o\u0000\u0000<\u000e\u0001\u0000\u0000\u0000"+
		"=>\u0005e\u0000\u0000>?\u0005n\u0000\u0000?@\u0005d\u0000\u0000@\u0010"+
		"\u0001\u0000\u0000\u0000AB\u0005:\u0000\u0000BC\u0005=\u0000\u0000C\u0012"+
		"\u0001\u0000\u0000\u0000DE\u0005[\u0000\u0000E\u0014\u0001\u0000\u0000"+
		"\u0000FG\u0005]\u0000\u0000G\u0016\u0001\u0000\u0000\u0000HI\u0005(\u0000"+
		"\u0000I\u0018\u0001\u0000\u0000\u0000JK\u0005)\u0000\u0000K\u001a\u0001"+
		"\u0000\u0000\u0000LM\u0005;\u0000\u0000M\u001c\u0001\u0000\u0000\u0000"+
		"NP\u0007\u0002\u0000\u0000ON\u0001\u0000\u0000\u0000PQ\u0001\u0000\u0000"+
		"\u0000QO\u0001\u0000\u0000\u0000QR\u0001\u0000\u0000\u0000R\u001e\u0001"+
		"\u0000\u0000\u0000SU\u0007\u0003\u0000\u0000TS\u0001\u0000\u0000\u0000"+
		"UV\u0001\u0000\u0000\u0000VT\u0001\u0000\u0000\u0000VW\u0001\u0000\u0000"+
		"\u0000W \u0001\u0000\u0000\u0000XZ\u0007\u0004\u0000\u0000YX\u0001\u0000"+
		"\u0000\u0000Z[\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000[\\\u0001"+
		"\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000]^\u0006\u0010\u0000\u0000"+
		"^\"\u0001\u0000\u0000\u0000\u0004\u0000QV[\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}