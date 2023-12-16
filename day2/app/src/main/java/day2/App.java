package day2;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class App {
    private static final int MAX_RED = 12;
    private static final int MAX_GREEN = 13;
    private static final int MAX_BLUE = 14;

    private static final String RED = "red";
    private static final String GREEN = "green";
    private static final String BLUE = "blue";

    public static int getGameNumber(String gameStr) {
        String[] splitGameStr = gameStr.split(":");
        String gameHeader = splitGameStr[0];

        String gameNumberStr = gameHeader.split(" ")[1];
        return Integer.parseInt(gameNumberStr);
    }

    public static boolean isValidGame(String gameRound) {
        String[] round = gameRound.split(", ");

        int countRed = 0;
        int countGreen = 0;
        int countBlue = 0;

        for (String balls : round) {
            String[] splitBalls = balls.split(" ");
            int ballCount = Integer.parseInt(splitBalls[0]);
            String ballColor = splitBalls[1];

            switch (ballColor) {
                case App.RED:
                    countRed += ballCount;
                    break;
                case App.GREEN:
                    countGreen += ballCount;
                    break;
                case App.BLUE:
                    countBlue += ballCount;
                    break;
                default:
                    break;
            }
        }

        return (
            countRed <=     App.MAX_RED &&
            countGreen <=   App.MAX_GREEN &&
            countBlue <=    App.MAX_BLUE
        );
    }

    public static int checkGame(String gameStr) {
        int gameNumber = getGameNumber(gameStr);

        String allGames = gameStr.split(": ")[1];
        String[] games = allGames.split("; ");

        for (String game : games) {
            boolean validGame = isValidGame(game);
            if (!validGame) {
                return -1;
            }
        }

        return gameNumber;
    }

    public static boolean isNumeric(String str) { 
        try {  
            Double.parseDouble(str);  
            return true;
        } catch(NumberFormatException e){  
            return false;  
        }  
    }

    public static int getMaxCount(String[] splitColor) {
        int maxCount = 0;
        for (String split : splitColor) {

            String[] spaceSplit = split.split(" ");
            String end = spaceSplit[spaceSplit.length - 1];
            if (isNumeric(end)) {
                int numEnd = Integer.parseInt(end);
                maxCount = (maxCount < numEnd) ? numEnd : maxCount;
            }
        }

        return maxCount;
    }

    public static int getPowerOfGame(String gameStr) {
        String[] splitGame = gameStr.split(": ");
        String gameHeader = splitGame[0];
        String gameRounds = splitGame[1];
        
        String[] splitRed = gameRounds.split(" " + App.RED);
        String[] splitGreen = gameRounds.split(" " + App.GREEN);
        String[] splitBlue = gameRounds.split(" " + App.BLUE);

        int maxRed = getMaxCount(splitRed);
        int maxGreen = getMaxCount(splitGreen);
        int maxBlue = getMaxCount(splitBlue);

        System.out.println(gameHeader + " " + maxRed + " " + maxGreen + " " + maxBlue);

        return maxRed * maxGreen * maxBlue;
    }

    public static void main(String[] args) {
        File input = new File("src/main/java/day2/day2.txt");
        try {
            Scanner scanner = new Scanner(input);

            int gameSum = 0;

            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                gameSum += getPowerOfGame(line);
            }

            System.out.println("Total: " + gameSum);

            scanner.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
    }
}
