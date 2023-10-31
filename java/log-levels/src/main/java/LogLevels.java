public class LogLevels {
    
    public static String message(String logLine) {
        return logLine.split(":")[1].trim();
        //throw new UnsupportedOperationException("Please implement the (static) LogLine.message() method");
    }

    public static String logLevel(String logLine) {
        return logLine.split(":")[0].substring(1).replaceAll("]","").toLowerCase();
        //throw new UnsupportedOperationException("Please implement the (static) LogLine.logLevel() method");
    }

    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
        //throw new UnsupportedOperationException("Please implement the (static) LogLine.reformat() method");
    }
}
