package cricinfo;

import cricinfo.entity.Ball;
import cricinfo.entity.Player;
import cricinfo.entity.Wicket;
import cricinfo.enums.ExtraType;
import cricinfo.enums.PlayerRole;
import cricinfo.enums.WicketType;

public final class CommentaryAttributionTest {
    public static void main(String[] args) {
        Player batter = new Player("batter", "Batter Name", PlayerRole.BATSMAN);
        Player bowler = new Player("bowler", "Bowler Name", PlayerRole.BOWLER);
        boolean sawBowler = false;
        for (int i = 0; i < 100; i++) {
            Ball noBall = new Ball.BallBuilder().bowledBy(bowler).facedBy(batter)
                    .withExtraType(ExtraType.NO_BALL).build();
            String noBallCommentary = noBall.getCommentary();
            if (noBallCommentary.contains("Batter Name")) {
                throw new AssertionError("no-ball overstep was attributed to the batter");
            }
            sawBowler |= noBallCommentary.contains("Bowler Name");
        }
        if (!sawBowler) {
            throw new AssertionError("no-ball templates never attributed overstepping to the bowler");
        }

        Wicket bowled = new Wicket.Builder(WicketType.BOWLED, batter).build();
        for (int i = 0; i < 100; i++) {
            Ball wicketBall = new Ball.BallBuilder().bowledBy(bowler).facedBy(batter)
                    .withWicket(bowled).build();
            if (!wicketBall.getCommentary().contains("Bowler Name")) {
                throw new AssertionError("bowled commentary was not attributed to the bowler");
            }
        }
    }
}
